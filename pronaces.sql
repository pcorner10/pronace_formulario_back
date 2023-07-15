-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS logs_id_seq;

-- Table Definition
CREATE TABLE "public"."logs" (
    "id" int8 NOT NULL DEFAULT nextval('logs_id_seq'::regclass),
    "user_id" int8 NOT NULL,
    "description" varchar(255) NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table0_id_seq;

-- Table Definition
CREATE TABLE "public"."table0" (
    "id" int8 NOT NULL DEFAULT nextval('table0_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "sexo" text,
    "edad" int8,
    "parentesco" text,
    "escolaridad" text,
    "trabaja" text,
    "horas_por_dia" int8,
    "dias_por_semana" int8,
    "ocupacion" text,
    "fumador" text,
    "tiempo_ciudad" int8,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    CONSTRAINT "fk_table0_encuestador" FOREIGN KEY ("encuestador_id") REFERENCES "public"."users"("id"),
    CONSTRAINT "fk_table0_zona" FOREIGN KEY ("zona_id") REFERENCES "public"."zonas"("id"),
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table1_id_seq;

-- Table Definition
CREATE TABLE "public"."table1" (
    "id" int8 NOT NULL DEFAULT nextval('table1_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "tiene_electricidad" text,
    "fuentes_agua" text,
    "tiene_gas" text,
    "eliminacion_desechos" text,
    "material_techo" text,
    "material_piso" text,
    "material_paredes" text,
    "num_cuartos" int8,
    "olores_desagradables" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table10_id_seq;

-- Table Definition
CREATE TABLE "public"."table10" (
    "id" int8 NOT NULL DEFAULT nextval('table10_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "discapacidad_referida" text,
    "discapacidad_recabada" text,
    "tipo_condicion" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table10_1_id_seq;

-- Table Definition
CREATE TABLE "public"."table10_1" (
    "id" int8 NOT NULL DEFAULT nextval('table10_1_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "si_donde" text,
    "no_porque" text,
    "tiene_certificado" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table11_id_seq;

-- Table Definition
CREATE TABLE "public"."table11" (
    "id" int8 NOT NULL DEFAULT nextval('table11_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "nombre_farmaco" text,
    "es_prescrito" text,
    "motivo_referido" text,
    "motivo_recabado" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table12_id_seq;

-- Table Definition
CREATE TABLE "public"."table12" (
    "id" int8 NOT NULL DEFAULT nextval('table12_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "hay_problema" text,
    "que_problema" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table13_id_seq;

-- Table Definition
CREATE TABLE "public"."table13" (
    "id" int8 NOT NULL DEFAULT nextval('table13_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "hay_contaminacion" text,
    "cual_es" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table14_id_seq;

-- Table Definition
CREATE TABLE "public"."table14" (
    "id" int8 NOT NULL DEFAULT nextval('table14_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "tiempo" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table2_id_seq;

-- Table Definition
CREATE TABLE "public"."table2" (
    "id" int8 NOT NULL DEFAULT nextval('table2_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "llave_interior" text,
    "garrafon" text,
    "llave_comunitaria" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table3_id_seq;

-- Table Definition
CREATE TABLE "public"."table3" (
    "id" int8 NOT NULL DEFAULT nextval('table3_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "imss" text,
    "issste" text,
    "seguro_popular" text,
    "privado" text,
    "ninguno" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table4_id_seq;

-- Table Definition
CREATE TABLE "public"."table4" (
    "id" int8 NOT NULL DEFAULT nextval('table4_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "enfermedad_referida" text,
    "enfermedad_recabada" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table5_id_seq;

-- Table Definition
CREATE TABLE "public"."table5" (
    "id" int8 NOT NULL DEFAULT nextval('table5_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "enfermedad_referida" text,
    "enfermedad_recabada" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table6_id_seq;

-- Table Definition
CREATE TABLE "public"."table6" (
    "id" int8 NOT NULL DEFAULT nextval('table6_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "sexo" text,
    "edad_fallecimiento" int8,
    "año_fallecimiento" int8,
    "causa_referida" text,
    "causa_recabada" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table7_id_seq;

-- Table Definition
CREATE TABLE "public"."table7" (
    "id" int8 NOT NULL DEFAULT nextval('table7_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "sexo" text,
    "edad_detection" int8,
    "año_detection" int8,
    "tipo_referido" text,
    "tipo_recabado" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table8_id_seq;

-- Table Definition
CREATE TABLE "public"."table8" (
    "id" int8 NOT NULL DEFAULT nextval('table8_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "año_nacimiento_perdida" int8,
    "en_curso" text,
    "tipo_parto" text,
    "tuvo_complicaciones" text,
    "complicacion_referida" text,
    "complicacion_recabada" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table8_1_id_seq;

-- Table Definition
CREATE TABLE "public"."table8_1" (
    "id" int8 NOT NULL DEFAULT nextval('table8_1_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "año_perdida" int8,
    "trimestre" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table9_id_seq;

-- Table Definition
CREATE TABLE "public"."table9" (
    "id" int8 NOT NULL DEFAULT nextval('table9_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "bajo_peso" bool,
    "prematuro" bool,
    "malformacion" bool,
    "malformacion_referida" text,
    "malformacion_recabada" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS table9_1_id_seq;

-- Table Definition
CREATE TABLE "public"."table9_1" (
    "id" int8 NOT NULL DEFAULT nextval('table9_1_id_seq'::regclass),
    "encuestador_id" int8,
    "zona_id" int8,
    "año" int8,
    "problema_referido" text,
    "problema_recabado" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "user_name" text NOT NULL,
    "password" text NOT NULL,
    "role" text NOT NULL,
    "email" text NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS zonas_id_seq;

-- Table Definition
CREATE TABLE "public"."zonas" (
    "id" int8 NOT NULL DEFAULT nextval('zonas_id_seq'::regclass),
    "municipio" text,
    "localidad" text,
    "ageb" text,
    "manzana" text,
    "lote" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

