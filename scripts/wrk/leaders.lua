wrk.host = "localhost:10000"
wrk.path = "/shitlist.v1.ShitlistService/Leaders"
wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"
wrk.body = "{}"