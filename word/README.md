# i18n-town

Let local people improve i18n data, instead of counting on translation company.

架構是
-> api router (由 gin 協助)
-> controller (主要驗證接收參數，傳遞資料)
-> service 與資料交互，或是通訊外部訊息, e.g grpc communicate with admin module
-> repo (處理資料，cache 也在這，也不會因為 DB 更換影響到上層)
-> model (SQL adapter，一開始使用 memory map 充當 db，後來改成 psql)
