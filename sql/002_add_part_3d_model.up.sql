CREATE TABLE IF NOT EXISTS part_3d_models (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    part_id UUID NOT NULL REFERENCES parts(id),

    url TEXT NOT NULL,

    file_name TEXT NOT NULL,
    file_size BIGINT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_part_3d_models_part_id ON part_3d_models(part_id);