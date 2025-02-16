-- Table: short_links
CREATE TABLE short_links
(
    id         INTEGER GENERATED BY DEFAULT AS IDENTITY
        CONSTRAINT pk_short_links PRIMARY KEY,
    url        TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_short_links_url ON short_links (url);