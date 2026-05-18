// Use modern DBML syntax
// Project: Edv (Learning Management System)

Enum material_type {
video
pdf
ppt
other
}

Enum source_type {
material
assignment
feed
submission
comment
}

Enum owner_type {
user
material
assignment
feed
submission
comment
school
system
}

Enum class_role {
teacher
student
}

Enum status_progress {
not_started
completed
}

Table schools {
sch_id uuid [pk, default: `gen_random_uuid()`]
sch_name varchar(150)
sch_code varchar(50) [unique]
sch_address text
sch_email text
sch_phone text
sch_website text
sch_logo uuid
created_at timestamp [default: `now()`]
updated_at timestamp [default: `now()`]
deleted_at timestamp
}

Table academic_years {
acy_id uuid [pk, default: `gen_random_uuid()`]
acy_sch_id uuid [ref: > schools.sch_id]
acy_name varchar(20)
is_active boolean [default: false]
created_at timestamp [default: `now()`]

indexes {
(acy_sch_id, acy_name) [unique]
}
}

Table terms {
trm_id uuid [pk, default: `gen_random_uuid()`]
trm_acy_id uuid [ref: > academic_years.acy_id]
trm_name varchar(10)
is_active boolean [default: false]
created_at timestamp [default: `now()`]

indexes {
(trm_acy_id, trm_name) [unique]
}
}

Table users {
usr_id uuid [pk, default: `gen_random_uuid()`]
usr_nama_lengkap varchar(150)
usr_email varchar(150) [not null]
usr_password varchar(255)
is_active boolean [default: true]
created_at timestamp [default: `now()`]
updated_at timestamp [default: `now()`]
deleted_at timestamp

indexes {
(usr_email, deleted_at) [unique]
}
}

Table school_users {
scu_id uuid [pk, default: `gen_random_uuid()`]
scu_usr_id uuid [ref: > users.usr_id]
scu_sch_id uuid [ref: > schools.sch_id]
created_at timestamp [default: `now()`]

indexes {
(scu_usr_id, scu_sch_id) [unique]
}
}

Table roles {
rol_id uuid [pk, default: `gen_random_uuid()`]
rol_name varchar(50)
created_at timestamp [default: `now()`]
}

Table user_roles {
urol_id uuid [pk, default: `gen_random_uuid()`]
urol_scu_id uuid [ref: > school_users.scu_id]
urol_rol_id uuid [ref: > roles.rol_id]
created_at timestamp [default: `now()`]

indexes {
(urol_scu_id, urol_rol_id) [unique]
}
}

Table medias {
med_id uuid [pk, default: `gen_random_uuid()`]
med_sch_id uuid [ref: > schools.sch_id]
med_name varchar(255)
med_file_size bigint
med_mime_type varchar(100)
med_storage_path text
med_file_url text
med_thumbnail_url text
is_public boolean [default: true]
med_owner_type owner_type
med_owner_id uuid
created_at timestamp [default: `now()`]
deleted_at timestamp
}

// Relationship for school logo
Ref: schools.sch_logo > medias.med_id

Table attachments {
att_id uuid [pk, default: `gen_random_uuid()`]
att_sch_id uuid [ref: > schools.sch_id]
att_source_id uuid
att_source_type source_type
att_med_id uuid [ref: > medias.med_id]
created_at timestamp [default: `now()`]
}

Table subjects {
sub_id uuid [pk, default: `gen_random_uuid()`]
sub_sch_id uuid [ref: > schools.sch_id]
sub_name varchar(100)
sub_code varchar(20)
created_at timestamp [default: `now()`]

indexes {
(sub_sch_id, sub_code) [unique]
}
}

Table classes {
cls_id uuid [pk, default: `gen_random_uuid()`]
cls_sch_id uuid [ref: > schools.sch_id]
cls_trm_id uuid [ref: > terms.trm_id]
cls_code varchar(20)
cls_title varchar(150)
cls_desc text
created_by uuid [ref: > users.usr_id]
is_active boolean [default: true]
created_at timestamp [default: `now()`]
updated_at timestamp [default: `now()`]
deleted_at timestamp

indexes {
(cls_sch_id, cls_code, cls_trm_id) [unique]
}
}

//subject_classes untuk kelas per mata pelajaran
Table subject_classes {
scl_id uuid [pk, default: `gen_random_uuid()`]
scl_cls_id uuid [ref: > classes.cls_id]
scl_sub_id uuid [ref: > subjects.sub_id]
scl_scu_id uuid [ref: > school_users.scu_id]

indexes {
(scl_cls_id, scl_sub_id, scl_scu_id) [unique]
}
}

//enrollments untuk kelas (misal 12 IPA)
Table enrollments {
enr_id uuid [pk, default: `gen_random_uuid()`]
enr_sch_id uuid [ref: > schools.sch_id]
enr_scu_id uuid [ref: > school_users.scu_id]
enr_cls_id uuid [ref: > classes.cls_id]
enr_role class_role
joined_at timestamp [default: `now()`]

indexes {
(enr_scu_id, enr_cls_id) [unique]
}
}

Table materials {
mat_id uuid [pk, default: `gen_random_uuid()`]
mat_sch_id uuid [ref: > schools.sch_id]

// Ubah dari cls_id menjadi scl_id
mat_scl_id uuid [ref: > subject_classes.scl_id]

mat_title varchar(150)
mat_desc text
mat_types material_type
created_by uuid [ref: > users.usr_id]
created_at timestamp [default: `now()`]
updated_at timestamp [default: `now()`]
deleted_at timestamp
}

Table material_progress {
map_id uuid [pk, default: `gen_random_uuid()`]
map_usr_id uuid [ref: > users.usr_id]
map_mat_id uuid [ref: > materials.mat_id]
map_status status_progress
last_opened_at timestamp

indexes {
(map_usr_id, map_mat_id) [unique]
}
}

Table feeds {
fds_id uuid [pk, default: `gen_random_uuid()`]
fds_sch_id uuid [ref: > schools.sch_id]
fds_cls_id uuid [ref: > classes.cls_id]
fds_content text
created_by uuid [ref: > users.usr_id]
created_at timestamp [default: `now()`]
deleted_at timestamp
}

Table comments {
cmn_id uuid [pk, default: `gen_random_uuid()`]
cmn_sch_id uuid [ref: > schools.sch_id]
cmn_source_type source_type
cmn_source_id uuid
cmn_usr_id uuid [ref: > users.usr_id]
cmn_content text
created_at timestamp [default: `now()`]
deleted_at timestamp
}

Table assignment_categories {
asc_id uuid [pk, default: `gen_random_uuid()`]
asc_sch_id uuid [ref: > schools.sch_id]
asc_name varchar(50)
created_at timestamp [default: `now()`]

indexes {
(asc_sch_id, asc_name) [unique]
}
}

Table assignments {
asg_id uuid [pk, default: `gen_random_uuid()`]
asg_sch_id uuid [ref: > schools.sch_id]
asg_scl_id uuid [ref: > subject_classes.scl_id]
asg_asc_id uuid [ref: > assignment_categories.asc_id]
asg_title varchar(150)
asg_desc text
asg_deadline timestamp
asg_allowed_late bool
created_by uuid [ref: > users.usr_id]
created_at timestamp [default: `now()`]
updated_at timestamp [default: `now()`]
deleted_at timestamp
}

Table submissions {
sbm_id uuid [pk, default: `gen_random_uuid()`]
sbm_sch_id uuid [ref: > schools.sch_id]
sbm_asg_id uuid [ref: > assignments.asg_id]
sbm_usr_id uuid [ref: > users.usr_id]
submitted_at timestamp [default: `now()`]
deleted_at timestamp

indexes {
(sbm_asg_id, sbm_usr_id) [unique]
}
}

Table assessments {
asm_id uuid [pk, default: `gen_random_uuid()`]
asm_sbm_id uuid [ref: > submissions.sbm_id]
asm_score decimal(5,2)
asm_feedback text
assessed_by uuid [ref: > users.usr_id]
assessed_at timestamp [default: `now()`]
}

Table assessments_weights {
asw_id uuid [pk, default: `gen_random_uuid()`]
asw_sub_id uuid [ref: > subjects.sub_id]
asw_asc_id uuid [ref: > assignment_categories.asc_id]
asw_weight decimal(5,2)

indexes {
(asw_sub_id, asw_asc_id) [unique]
}
}

Table logs {
log_id uuid [pk, default: `gen_random_uuid()`]
log_sch_id uuid [ref: > schools.sch_id]
log_usr_id uuid [ref: > users.usr_id]
log_action varchar(150)
log_metadata jsonb
created_at timestamp [default: `now()`]
}

Table notifications {
ntf_id uuid [pk, default: `gen_random_uuid()`]
ntf_usr_id uuid [ref: > users.usr_id]
ntf_type varchar(50)
ntf_title varchar(255)
ntf_message text
ntf_link text
ntf_related_id uuid
is_read boolean [default: false]
created_at timestamp [default: `now()`]

indexes {
(ntf_usr_id, is_read, created_at) [name: 'idx_notifications_user']
}
}
