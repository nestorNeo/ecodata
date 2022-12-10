import re
import json
import requests

consulta_web = requests.get("https://aire.jalisco.gob.mx/porestacion")
consulta_web.raise_for_status()

datos = consulta_web.text #io.open("poresta.html", mode="r", encoding="utf-8").read()

pattern_start = re.compile(r"w3DisplayData\(\"Hoy\"", re.MULTILINE)
pattern_end = re.compile(r"w3DisplayData\(\"Ayer\"", re.MULTILINE)

inicio = pattern_start.search(datos)
fin = pattern_end.search(datos)

print(inicio)
print(fin)
datos_cortados = datos[inicio.end():fin.start()]

pattern_start = re.compile(r"\{", re.MULTILINE)
pattern_end = re.compile(r"\)", re.MULTILINE)

inicio = pattern_start.search(datos_cortados)
fin = pattern_end.search(datos_cortados)
print(inicio)
print(fin)

datos_cortados = datos_cortados[inicio.start():fin.start()]
print(datos_cortados)

contaminacion = json.loads(datos_cortados)
print(contaminacion["MyDataBinding"])


from db import get_database

conn = get_database()
realdb = conn.get_database("contaminacion")
collection = realdb.get_collection("aire")

for entry in contaminacion["MyDataBinding"]:
    collection.insert_one(entry)