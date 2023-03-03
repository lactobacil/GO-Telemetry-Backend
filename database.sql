create table site(
	id SERIAL primary key,
	site SMALLINT,
	site_name  VARCHAR(50)
);

create table supervisor_code(
	id SERIAL primary key,
	supervisor_code VARCHAR(50)
);

create table plant_area(
	id SERIAL primary key,
	plant_area VARCHAR(50),
	site_id smallint references site(id)
);

create table cell_number(
	id 	SERIAL primary key,
	cell_number smallint,
	cell_name VARCHAR(50),
	supervisor_code_id smallint references supervisor_code(id),
	plant_area_id smallint  references plant_area(id) 
);

create table work_centre(
	id SERIAL primary key,
	work_centre VARCHAR(50),
	cell_number_id smallint  references cell_number(id),
	work_centre_display_name VARCHAR(50)
);

create table sfm(
	id 	SERIAL primary key,
	sfm VARCHAR(50),
	sfm_display VARCHAR(50),
	sfm_description VARCHAR(50),
	machine_model VARCHAR,
	install_date TIMESTAMP,
	work_centre_id INT  references work_centre(id),
	bottleneck_status VARCHAR(50),
	sfm_importance VARCHAR(50),
	sfm_latest_status_timestamp VARCHAR(50),
	andon_link VARCHAR(100)
);  

create table andon_transaction(
	id 		SERIAL primary key,
	sfm_id 	smallint references sfm(id),
	andon_status VARCHAR(50),
	andon_complain_start_time TIMESTAMP,
	andon_complain_finish_time TIMESTAMP,
	andon_name VARCHAR(50)
);