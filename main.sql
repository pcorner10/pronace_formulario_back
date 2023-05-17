
CREATE TABLE "encuestador" (
"id" integer PRIMARY KEY,
"username" varchar,
"role" varchar,
"created_at" timestamp
);

CREATE TABLE "encuestado" (
"id" integer PRIMARY KEY,
"encuestadorId" integer REFERENCES "encuestador" ("id"),
"Sexo" varchar,
"Edad" integer,
"Fuma" bool,
"TiempoResidencia" integer,
"escolaridadID" integer REFERENCES "escolaridad" ("id"),
"ocupacionID" integer REFERENCES "ocupacion" ("id"),
"zonaID" integer REFERENCES "zonaEncuestado" ("id")
);

CREATE TABLE "vivienda" (
"id" integer PRIMARY KEY,
"Electricidad" bool,
"FuenteAgua" varchar,
"Gas" bool,
"WCDentro" bool,
"NoHabitaciones" integer,
"lugarDeAtencionId" integer REFERENCES "lugarAtencion" ("id"),
"oloresId" integer REFERENCES "olores" ("id"),
"AguasNegrasId" integer REFERENCES "aguasNegras" ("id"),
"TechoId" integer REFERENCES "materialTecho" ("id"),
"ParedesId" integer REFERENCES "materialParedes" ("id"),
"PisoId" integer REFERENCES "materialPiso" ("id")
);

CREATE TABLE "problemaSalud" (
"id" integer PRIMARY KEY,
"viviendaId" integer REFERENCES "vivienda" ("id"),
"encuestadoID" integer REFERENCES "encuestado" ("id"),
"efectacion" bool,
"mas12meses" bool,
"referido" varchar,
"recabado" varchar,
"cieId" integer REFERENCES "cie" ("id")
);

CREATE TABLE "cie" (
"id" integer PRIMARY KEY,
"nombre" integer
);

CREATE TABLE "lugarAtencion" (
"id" integer PRIMARY KEY,
"Nombre" varchar
);

CREATE TABLE "olores" (
"id" integer PRIMARY KEY,
"encuestadoId" integer REFERENCES "encuestado" ("id"),
"olor" bool,
"Descripcion" varchar
);

CREATE TABLE "aguasNegras" (
"id" integer PRIMARY KEY,
"Nombre" varchar
);

CREATE TABLE "materialTecho" (
"id" integer PRIMARY KEY,
"Nombre" varchar
);

CREATE TABLE "materialParedes" (
"id" integer PRIMARY KEY,
"Nombre" varchar
);

CREATE TABLE "materialPiso" (
"id" integer PRIMARY KEY,
"Nombre" varchar
);

CREATE TABLE "zonaEncuestado" (
"id" integer PRIMARY KEY,
"Calle" varchar,
"Localidad" varchar,
"Manzana" varchar,
"Ageb" varchar
);

CREATE TABLE "escolaridad" (
"id" integer PRIMARY KEY,
"nivelEscolar" varchar,
"Trunca" bool,
"NoEscolarizado" bool,
"NoCorresponde" bool
);

CREATE TABLE "ocupacion" (
"id" integer PRIMARY KEY,
"nombreOcupacion" varchar,
"TrabajoRemunerado" bool,
"HorasPorDia" integer,
"DiasPorSemana" integer,
"NoCorresponde" bool
);

CREATE TABLE "parentesco" (
"id" integer PRIMARY KEY,
"descripcionParentesco" varchar
);

ALTER TABLE "encuestado" ADD FOREIGN KEY ("encuestadorId") REFERENCES "encuestador" ("id");
ALTER TABLE "encuestado" ADD FOREIGN KEY ("escolaridadID") REFERENCES "escolaridad" ("id");
ALTER TABLE "encuestado" ADD FOREIGN KEY ("ocupacionID") REFERENCES "ocupacion" ("id");
ALTER TABLE "encuestado" ADD FOREIGN KEY ("zonaID") REFERENCES "zonaEncuestado" ("id");

ALTER TABLE "vivienda" ADD FOREIGN KEY ("lugarDeAtencionId") REFERENCES "lugarAtencion" ("id");
ALTER TABLE "vivienda" ADD FOREIGN KEY ("oloresId") REFERENCES "olores" ("id");
ALTER TABLE "vivienda" ADD FOREIGN KEY ("AguasNegrasId") REFERENCES "aguasNegras" ("id");
ALTER TABLE "vivienda" ADD FOREIGN KEY ("TechoId") REFERENCES "materialTecho" ("id");
ALTER TABLE "vivienda" ADD FOREIGN KEY ("ParedesId") REFERENCES "materialParedes" ("id");
ALTER TABLE "vivienda" ADD FOREIGN KEY ("PisoId") REFERENCES "materialPiso" ("id");

ALTER TABLE "problemaSalud" ADD FOREIGN KEY ("viviendaId") REFERENCES "vivienda" ("id");
ALTER TABLE "problemaSalud" ADD FOREIGN KEY ("encuestadoID") REFERENCES "encuestado" ("id");
ALTER TABLE "problemaSalud" ADD FOREIGN KEY ("cieId") REFERENCES "cie" ("id");

ALTER TABLE "olores" ADD FOREIGN KEY ("encuestadoId") REFERENCES "encuestado" ("id");