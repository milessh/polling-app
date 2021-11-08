CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE polling_strategy (
  id INTEGER PRIMARY KEY,
  strategy TEXT NOT NULL,
  display_name TEXT NOT NULL,
  description TEXT NOT NULL
);
COMMENT ON TABLE polling_strategy IS 'Describes the different voting and tallying strategies to employ';

GRANT ALL ON TABLE polling_strategy TO "polling-app-user";

INSERT INTO polling_strategy (id, strategy, display_name, description) VALUES (1, 'fptp', 'First Past the Post', 'Single vote, highest tally wins');
INSERT INTO polling_strategy (id, strategy, display_name, description) VALUES (2, 'weighted', 'Ranked, Weighted', 'Voting options weighted by position');
INSERT INTO polling_strategy (id, strategy, display_name, description) VALUES (3, 'runoff', 'Ranked, Run-off', 'Voting options weighted, removing lowest tallied option until majority');
INSERT INTO polling_strategy (id, strategy, display_name, description) VALUES (4, 'rating', 'Multiple Rating', 'Rate each option, tallying across all votes and candidates');


CREATE TABLE poll (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  strategy_id INTEGER REFERENCES polling_strategy(id) NOT NULL,
  poll_start TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  poll_end TIMESTAMP WITH TIME ZONE,
  user_hash TEXT NOT NULL,
  options TEXT NOT NULL,
  results TEXT
);
COMMENT ON TABLE poll IS 'Represents a poll';
COMMENT ON COLUMN poll.user_hash IS 'Hashed combination of client-ip and user-agent. We do not want to store personally identifiable data';
-- options and results are not intended to be directly queriable
COMMENT ON COLUMN poll.options IS 'A raw JSON string containing the vote/poll options';
COMMENT ON COLUMN poll.results IS 'A raw JSON string containing the results of the closed poll';

GRANT ALL ON TABLE poll TO "polling-app-user";
