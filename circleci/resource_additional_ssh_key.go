package circleci

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/crypto/ssh"
)

func resourceAdditionalSSHKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAdditionalSSHKeyCreate,
		ReadContext:   resourceAdditionalSSHKeyRead,
		DeleteContext: resourceAdditionalSSHKeyDelete,
		Schema: map[string]*schema.Schema{
			"project_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: assureSlugHasValidVCS,
			},
			"host_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAdditionalSSHKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	slug := d.Get("project_slug").(string)
	hostname := d.Get("host_name").(string)

	key := d.Get("private_key").(string)
	parsedKey, keyError := ssh.ParsePrivateKey([]byte(key))

	if keyError != nil {
		return diag.Errorf("Received error '%s' while trying to parse private_key", keyError.Error())
	}

	code, body, err := postToLegacyEndpoint(ctx, slug, hostname, key, m)

	if err != nil {
		return diag.FromErr(err)
	}

	if code != http.StatusOK {
		diag.Errorf("received status code %d with body %s when trying to post ssh key", code, string(body))
	}

	fingerprint := ssh.FingerprintLegacyMD5(parsedKey.PublicKey())
	id := fmt.Sprintf("%s/%s", slug, fingerprint)

	d.SetId(id)
	d.Set("fingerprint", fingerprint)

	return nil
}

func resourceAdditionalSSHKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// CircleCI's api does not offer any way to read additional ssh keys
	return nil
}

func resourceAdditionalSSHKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	slug := d.Get("project_slug").(string)
	hostname := d.Get("host_name").(string)
	fingerprint := d.Get("fingerprint").(string)

	code, body, err := deleteFromLegacyEndpoint(ctx, slug, hostname, fingerprint, m)

	if err != nil {
		return diag.FromErr(err)
	}

	if code != http.StatusOK {
		diag.Errorf("received status code %d with body %s when trying to delete ssh key", code, string(body))
	}

	return nil
}

func postToLegacyEndpoint(ctx context.Context, slug, hostname, key string, m interface{}) (int, []byte, error) {

	payload, err := json.Marshal(map[string]string{
		"hostname":    hostname,
		"private_key": key,
	})

	if err != nil {
		return 0, []byte(""), err
	}

	return sendRequestToLegacyEndpoint(ctx, http.MethodPost, slug, payload, m)
}

func deleteFromLegacyEndpoint(ctx context.Context, slug, hostname, fingerprint string, m interface{}) (int, []byte, error) {

	payload, err := json.Marshal(map[string]string{
		"hostname":    hostname,
		"fingerprint": fingerprint,
	})

	if err != nil {
		return 0, []byte(""), err
	}

	return sendRequestToLegacyEndpoint(ctx, http.MethodDelete, slug, payload, m)
}

func sendRequestToLegacyEndpoint(ctx context.Context, verb, slug string, payload []byte, m interface{}) (statusCode int, body []byte, err error) {

	switch verb {
	case http.MethodPost:
		break
	case http.MethodDelete:
		break
	default:
		err = fmt.Errorf("only supports POST and DELETE, not %s", verb)
		return
	}

	url := fmt.Sprintf("https://circleci.com/api/v1.1/project/%s/ssh-key", slug)

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, verb, url, bytes.NewBuffer(payload))
	if err != nil {
		return
	}

	apiKey := m.(*providerContext).apiKey

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Circle-Token", apiKey)

	provider := m.(*providerContext)
	client := provider.circleCiClient.GetConfig().HTTPClient

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	statusCode = resp.StatusCode
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}
