# Details

Date : 2023-03-08 17:12:55

Directory c:\\Users\\Lietotājs\\Documents\\GitHub\\conspiracy\\backend-social-media

Total : 49 files,  2041 codes, 302 comments, 307 blanks, all 2650 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [backend-social-media/.github/workflows/go.yml](/backend-social-media/.github/workflows/go.yml) | YAML | 89 | 11 | 4 | 104 |
| [backend-social-media/.idea/backend-social-media.iml](/backend-social-media/.idea/backend-social-media.iml) | XML | 10 | 0 | 0 | 10 |
| [backend-social-media/.idea/modules.xml](/backend-social-media/.idea/modules.xml) | XML | 9 | 0 | 0 | 9 |
| [backend-social-media/.idea/vcs.xml](/backend-social-media/.idea/vcs.xml) | XML | 7 | 0 | 0 | 7 |
| [backend-social-media/.idea/workspace.xml](/backend-social-media/.idea/workspace.xml) | XML | 85 | 0 | 0 | 85 |
| [backend-social-media/README.md](/backend-social-media/README.md) | Markdown | 7 | 0 | 3 | 10 |
| [backend-social-media/functions/Base64ED.go](/backend-social-media/functions/Base64ED.go) | Go | 14 | 1 | 4 | 19 |
| [backend-social-media/functions/GetStringBetween.go](/backend-social-media/functions/GetStringBetween.go) | Go | 27 | 2 | 4 | 33 |
| [backend-social-media/functions/RemoveItemFromArrayByIndex.go](/backend-social-media/functions/RemoveItemFromArrayByIndex.go) | Go | 16 | 0 | 2 | 18 |
| [backend-social-media/functions/cookieStringtoJSON.go](/backend-social-media/functions/cookieStringtoJSON.go) | Go | 36 | 1 | 6 | 43 |
| [backend-social-media/functions/cooldown/cooldown.go](/backend-social-media/functions/cooldown/cooldown.go) | Go | 1 | 0 | 0 | 1 |
| [backend-social-media/functions/csrf/csrf.go](/backend-social-media/functions/csrf/csrf.go) | Go | 103 | 3 | 29 | 135 |
| [backend-social-media/functions/derefString.go](/backend-social-media/functions/derefString.go) | Go | 7 | 0 | 3 | 10 |
| [backend-social-media/functions/mongo/database.go](/backend-social-media/functions/mongo/database.go) | Go | 127 | 57 | 9 | 193 |
| [backend-social-media/functions/pointer.go](/backend-social-media/functions/pointer.go) | Go | 4 | 0 | 0 | 4 |
| [backend-social-media/functions/randString.go](/backend-social-media/functions/randString.go) | Go | 10 | 0 | 3 | 13 |
| [backend-social-media/functions/rateLimiter/rateLimiter.go](/backend-social-media/functions/rateLimiter/rateLimiter.go) | Go | 16 | 9 | 3 | 28 |
| [backend-social-media/functions/savePreviousURLtoSession.go](/backend-social-media/functions/savePreviousURLtoSession.go) | Go | 12 | 0 | 4 | 16 |
| [backend-social-media/functions/security/encrypt.go](/backend-social-media/functions/security/encrypt.go) | Go | 61 | 3 | 6 | 70 |
| [backend-social-media/functions/sessions/session.go](/backend-social-media/functions/sessions/session.go) | Go | 118 | 14 | 13 | 145 |
| [backend-social-media/functions/snowflake/postID.go](/backend-social-media/functions/snowflake/postID.go) | Go | 233 | 58 | 76 | 367 |
| [backend-social-media/functions/validXML.go](/backend-social-media/functions/validXML.go) | Go | 5 | 0 | 2 | 7 |
| [backend-social-media/functions/xss.go](/backend-social-media/functions/xss.go) | Go | 27 | 1 | 5 | 33 |
| [backend-social-media/go.mod](/backend-social-media/go.mod) | Go Module File | 49 | 0 | 4 | 53 |
| [backend-social-media/go.sum](/backend-social-media/go.sum) | Go Checksum File | 157 | 0 | 1 | 158 |
| [backend-social-media/main.go](/backend-social-media/main.go) | Go | 78 | 47 | 33 | 158 |
| [backend-social-media/routes/auth/get.go](/backend-social-media/routes/auth/get.go) | Go | 42 | 0 | 2 | 44 |
| [backend-social-media/routes/auth/user.go](/backend-social-media/routes/auth/user.go) | Go | 77 | 1 | 7 | 85 |
| [backend-social-media/routes/cookie/cookie.go](/backend-social-media/routes/cookie/cookie.go) | Go | 25 | 2 | 4 | 31 |
| [backend-social-media/routes/loader.go](/backend-social-media/routes/loader.go) | Go | 41 | 15 | 6 | 62 |
| [backend-social-media/routes/message/broadcast.go](/backend-social-media/routes/message/broadcast.go) | Go | 1 | 0 | 0 | 1 |
| [backend-social-media/routes/message/create.go](/backend-social-media/routes/message/create.go) | Go | 73 | 14 | 9 | 96 |
| [backend-social-media/routes/user/admin.go](/backend-social-media/routes/user/admin.go) | Go | 55 | 18 | 7 | 80 |
| [backend-social-media/routes/user/ban.go](/backend-social-media/routes/user/ban.go) | Go | 1 | 0 | 0 | 1 |
| [backend-social-media/routes/ws.go](/backend-social-media/routes/ws.go) | Go | 32 | 19 | 19 | 70 |
| [backend-social-media/sh/build.sh](/backend-social-media/sh/build.sh) | Shell Script | 2 | 0 | 1 | 3 |
| [backend-social-media/sh/echo.sh](/backend-social-media/sh/echo.sh) | Shell Script | 1 | 0 | 1 | 2 |
| [backend-social-media/sh/nohup_start.sh](/backend-social-media/sh/nohup_start.sh) | Shell Script | 4 | 0 | 1 | 5 |
| [backend-social-media/sh/pull.sh](/backend-social-media/sh/pull.sh) | Shell Script | 1 | 0 | 1 | 2 |
| [backend-social-media/sh/push.sh](/backend-social-media/sh/push.sh) | Shell Script | 1 | 0 | 2 | 3 |
| [backend-social-media/sh/start.sh](/backend-social-media/sh/start.sh) | Shell Script | 4 | 0 | 1 | 5 |
| [backend-social-media/sh/stop.sh](/backend-social-media/sh/stop.sh) | Shell Script | 1 | 0 | 1 | 2 |
| [backend-social-media/static/assets/index-08f9f5ff.js](/backend-social-media/static/assets/index-08f9f5ff.js) | JavaScript | 65 | 6 | 1 | 72 |
| [backend-social-media/static/assets/index-e3b202fe.css](/backend-social-media/static/assets/index-e3b202fe.css) | PostCSS | 1 | 0 | 1 | 2 |
| [backend-social-media/static/index.html](/backend-social-media/static/index.html) | HTML | 14 | 0 | 2 | 16 |
| [backend-social-media/static/vite.svg](/backend-social-media/static/vite.svg) | XML | 1 | 0 | 0 | 1 |
| [backend-social-media/structs/cookie.go](/backend-social-media/structs/cookie.go) | Go | 1 | 0 | 2 | 3 |
| [backend-social-media/structs/message.go](/backend-social-media/structs/message.go) | Go | 87 | 2 | 9 | 98 |
| [backend-social-media/structs/user.go](/backend-social-media/structs/user.go) | Go | 203 | 18 | 16 | 237 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)