pragma solidity>=0.4.24 <=0.7.0;
pragma experimental ABIEncoderV2;

import "./Table.sol";

contract OpenSource {
    TableFactory tableFactory;
    string constant REPO_TABLE = "repo";
    string constant USER_TABLE = "user";

    constructor() public {
        tableFactory = TableFactory(0x1001); //The fixed address is 0x1001 for TableFactory
        // the parameters of createTable are tableName,keyField,"vlaueFiled1,vlaueFiled2,vlaueFiled3,..."
        tableFactory.createTable(REPO_TABLE, "token_name", "owner,total_supply");
        tableFactory.createTable(USER_TABLE, "user", "token_name,balance");
    }

    //insert records
    function insert(string memory token_name, string memory owner, int256 total_supply)
    // function insert(string memory token_name, string memory owner, int256 total_supply, string memory user, int256 balance)
        public
        returns (int256)
    {
        Table repo_table = tableFactory.openTable(REPO_TABLE);
        Table user_table = tableFactory.openTable(USER_TABLE);

        Entry repo_entry = repo_table.newEntry();
        repo_entry.set("token_name", token_name);
        repo_entry.set("owner", owner);
        repo_entry.set("total_supply", total_supply);

        int256 count = repo_table.insert(token_name, repo_entry);
        // emit InsertResult(count);

        return count;
    }
}