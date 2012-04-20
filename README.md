# htdigest-go

Command line tool for managing htdigest files.


## Usage

    htdigest <htfile> <add|del> <realm> <user>

- `htfile`
    - Filename with htdigest records.
- `add|del`
    - Action to perform (add or delete user, realm pair).  
      Adding existing user will overwrite him.
- `realm`
    - Realm name, no comments needed, I hope.
- `user`
    - User name, also no comments needed.

### Examples

    # add user 'Magik' to 'Roorback' realm
    htdigest my_digest_file add Roorback Magik

    # add user 'Magik' to 'Team' realm
    htdigest my_digest_file add Team Magik

    # add user 'Shadow' to 'Team' realm
    htdigest my_digest_file add Team Shadow

    # change password for user 'Magik' in 'Roorback' realm
    htdigest my_digest_file add Roorback Magik

    # delete user 'Shadow' from 'Team' realm
    htdigest my_digest_file del Team Shadow    


## Installation

    go get github.com/MagikBSD/htdigest-go
