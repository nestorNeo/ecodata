### Instructions 
Users will use jwt to auth the rest api, 
so each user or device must have a token when security enable see 

```
.\security.exe -email test@ecodata.leofigy.xyz -key ..\..\security\id_rsa
TOKEN ------------------------
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZWNvZGF0YS5sZW9maWd5Lnh5eiIsImJpcnRoZGF5Ijo4Mjc0NTI4MDAsImV4cCI6MTY3NTEzMzc0MCwiaXNzIjoiRnJpZW5kcyBvZiBHbyJ9.kPlcVgOrRFl-WzpHEsAILimuTD_nYMuU_fBQYX35rHdy0z349lcV4T5wB3jcWZEsuEObKJFwgGaF_b3qGnI994C3Aqh7zdarw1xh4TwqakJpm1tOybG4CaHrZsl1ygp51madGhArbSM0Mfjm91-9Av6nJ5P0nKn1LxMsuhs1dFoJV6WCi8mMVMTNPhuPqyx8FCY2NhLGsXDxMelKsm0xXJd2XdMSVUWA_B2EcjnIJ25fX-nFQa2h3n1pgEDfngG5Zz2D3x8rrYHj9pCca5LVIjNL8CEPcsR2YU-dqY7h37-aCvTfZp2_50WW_xWV3H5BLzd2qLzZsWGpSx3xZ-mQmw
TOKEN ------------------------
map[birthday:8.274528e+08 email:test@ecodata.leofigy.xyz exp:1.67513374e+09 iss:Friends of Go]

```