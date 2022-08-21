-- This script was generated by a beta version of the ERD tool in pgAdmin 4.
-- Please log an issue at https://redmine.postgresql.org/projects/pgadmin4/issues/new if you find any bugs, including reproduction steps.
BEGIN;


CREATE TABLE public."blogPost"
(
    "blogPostId" integer NOT NULL,
    "blogPost" character varying(10000) NOT NULL,
    PRIMARY KEY ("blogPostId")
);

CREATE TABLE public."blogPreview"
(
    "blogId" integer NOT NULL,
    "blogName" character varying(40) NOT NULL,
    "blogDesc" character varying(150) NOT NULL,
    "blogKeyword1" character varying(25),
    "blogKeyword2" character varying(25),
    "blogKeyword3" character varying(25),
    "blogGenre" character varying(12) NOT NULL,
    "blogDateCreated" character varying(15) NOT NULL,
    "blogTimeToRead" integer NOT NULL,
    "blogPostId" integer NOT NULL,
    PRIMARY KEY ("blogId")
);

CREATE TABLE public.project
(
    "projectId" integer NOT NULL,
    "projectName" character varying(80) NOT NULL,
    "projectDescription" character varying(255) NOT NULL,
    "projectFeature1" character varying(140) NOT NULL,
    "projectFeature2" character varying(140) NOT NULL,
    "projectFeature3" character varying(140) NOT NULL,
    "projectGithub" character varying(200),
    "projectWebsite" character varying(200),
    PRIMARY KEY ("projectId")
);

CREATE TABLE public.workexperience
(
    "workId" integer NOT NULL,
    "companyName" character varying(50) NOT NULL,
    "position" character varying(50) NOT NULL,
    responsibility1 character varying(255) NOT NULL,
    responsibility2 character varying(255) NOT NULL,
    responsibility3 character varying(255) NOT NULL,
    "fromDate" character varying(25) NOT NULL,
    "toDate" character varying(25) NOT NULL,
    PRIMARY KEY ("workId")
);
END;