CREATE TABLE %s.transfers (
  from_stop_id char({{from_stop_id}}) NOT NULL,
  to_stop_id char({{to_stop_id}}) NOT NULL,
  transfer_type integer NOT NULL DEFAULT '0',
  min_transfer_time integer DEFAULT NULL --,
  -- PRIMARY KEY (from_stop_id,to_stop_id,transfer_type)
);
