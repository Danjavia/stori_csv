# Web version

- https://stori-csv.vercel.app/

# How To Run Backend (cmd) Folder

- Clone repository
- Set aws profile
- Intall serverless globally `yarn global add serverless`
- Run `make deploy` to generate cloudformation on AWS Cloud

## Enabled Services

- API_GATEWAY_URL/prod/transactions

This endpoint receives the file content of `txns.csv` in JSON format and transform to `Transaction Struct`
`cmd/api/transactions.go:15` and save data into `summary` table on Dynamo in AWS

- API_GATEWAY_URL/prod/send-email

This endpoint receives the `Transaction Struct` formatted to JSON data plus `userId`, `userEmail`, `artifactUrl` and send email with template located on AWS via `Amazon Simple Email Service` with custom template `STORI_TMPL` on `cmd/api/send-email.go:89` based on data via API

!(sample email)[https://i.ibb.co/xJMW39Q/Captura-de-pantalla-2024-01-05-a-la-s-2-14-26-a-m.png]

# How To Run Frontend (web) Folder

- copy file `.env.example` -> `.env`
- run `yarn start`
- open browser at (http://localhost:3000)
