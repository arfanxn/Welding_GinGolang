-- Create table
CREATE TABLE activities (
    id CHAR(26) PRIMARY KEY NOT NULL,

    causer_id VARCHAR(50),
    causer_type VARCHAR(50),
    causer_ip_address VARCHAR(45), -- if null then it is system activity

    action VARCHAR(50) NOT NULL,

    subject_id CHAR(26),
    subject_type VARCHAR(50),
    properties JSONB,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes
CREATE INDEX activities_causer_id_index ON activities (causer_id);
CREATE INDEX activities_causer_type_index ON activities (causer_type);
CREATE INDEX activities_ip_address_index ON activities (causer_ip_address);
CREATE INDEX activities_action_index ON activities (action);
CREATE INDEX activities_subject_id_index ON activities (subject_id);
CREATE INDEX activities_subject_type_index ON activities (subject_type);

