CREATE TABLE IF NOT EXISTS gebruiker (
     klant_id serial PRIMARY KEY,
     gebruikersnaam VARCHAR (50) UNIQUE NOT NULL,
    wachtwoord VARCHAR (520) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    region VARCHAR (255) NOT NULL,
    country VARCHAR (255) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS vraag (
     vraag_id serial PRIMARY KEY,
     vaknaam VARCHAR(20) NOT NULL,
    lesnaam VARCHAR(20) NOT NULL,
    punten_beschikbaar INT NOT NULL
    );

CREATE TABLE IF NOT EXISTS antwoord (
    resultaat INT,
    tijd_besteed INT,
    klant_id INT,
    vraag_id INT
);


CREATE TABLE IF NOT EXISTS studentenpas (
    pas_id serial PRIMARY KEY,
    klant_id INT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS cursus (
      cursus_id serial PRIMARY KEY,
      cursusnaam VARCHAR(20) NOT NULL,
    leraar_id INT NOT NULL
    );

CREATE TABLE IF NOT EXISTS leraar (
      leraar_id serial PRIMARY KEY,
      voornaam VARCHAR(20) NOT NULL,
    achternaam VARCHAR(20) NOT NULL,
    e_mail VARCHAR(20) UNIQUE NOT NULL
    );
