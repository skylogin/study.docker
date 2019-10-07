# Front-end
frontend_nginx - frontend_web

# Application
app_nginx - app_api

# DB
mysql_master (write) / mysql_slave (read)

---

# request flow
User -> frontend_nginx -> (proxy) -> frontend-web -> (http) -> app_nginx -> (proxy) -> app_api