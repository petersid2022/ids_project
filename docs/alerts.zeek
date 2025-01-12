@load base/protocols/http

const payload_threshold = 5000;
global body_buffer: table[string] of string &default = "";
global path_traversal_keywords = /ETC|PASSWD|USR|SHARE|\.\./i;
global code_injection_keywords = /IMG|ONERROR|SELECT|INSERT|UPDATE|DELETE|DROP|FROM|WHERE|UNION|--|#|;/i;

event http_request(c: connection, method: string, original_URI: string, unescaped_URI: string, version: string) {
  if (path_traversal_keywords in original_URI) {
    print fmt("[ALERT] Potential Path traversal detected. Host: %s:%d, URI: %s", c$id$orig_h, c$id$orig_p, original_URI);
  }
}

event http_entity_data(c: connection, is_orig: bool, length: count, data: string) {
  if (is_orig) {
    local conn_key = fmt("%s:%d->%s:%d", c$id$orig_h, c$id$orig_p, c$id$resp_h, c$id$resp_p);
    body_buffer[conn_key] += data;
  }
}

event http_end_entity(c: connection, is_orig: bool) {
  if (is_orig) {
    local conn_key = fmt("%s:%d->%s:%d", c$id$orig_h, c$id$orig_p, c$id$resp_h, c$id$resp_p);

    if (conn_key in body_buffer) {
      local full_body = body_buffer[conn_key];

      if (code_injection_keywords in full_body) {
        print fmt("[ALERT] Potential Code Injection detected. Host: %s:%d, Body: %s", c$id$orig_h, c$id$orig_p, full_body);
      }

      if (|full_body| > payload_threshold) {
        print fmt("[ALERT] Payload exceeds threshold. Host: %s:%d, Length: %s", c$id$orig_h, c$id$orig_p, |full_body|);
      }

      delete body_buffer[conn_key];
    }
  }
}
