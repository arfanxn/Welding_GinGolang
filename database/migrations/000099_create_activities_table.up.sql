-- Create enum types 
CREATE TYPE activity_action_enum AS ENUM (
    'users.register',
    'users.reset_password',
    'users.verify_email',
    'users.login',
    'users.logout',
    'users.me',
    'users.update_me_profile',
    'users.update_me_password',
    'users.index',
    'users.show',
    'users.store',
    'users.update',
    'users.toggle_activation',
    'users.destroy',

    'permissions.index',

    'roles.index',
    'roles.show',
    'roles.store',
    'roles.update',
    'roles.set_default',
    'roles.destroy',

    'codes.create_user_register_invitation',
    'codes.create_user_email_verification',
    'codes.create_user_reset_password'
);

CREATE TYPE activity_causer_type_enum AS ENUM (
    'user',
    'email',
    'guest',
    'system'
);

CREATE TYPE activity_subject_type_enum AS ENUM (
    'user',
    'permission',
    'role',
    'code'
);

-- Create table
CREATE TABLE activities (
    id CHAR(26) PRIMARY KEY NOT NULL,

    causer_id VARCHAR(50),
    causer_type activity_causer_type_enum,
    causer_ip_address VARCHAR(45), -- if null then it is system activity

    action activity_action_enum NOT NULL,

    subject_id CHAR(26),
    subject_type activity_subject_type_enum,
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

