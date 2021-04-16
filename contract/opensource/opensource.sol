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
        tableFactory.createTable(REPO_TABLE, "token_name", "owner,total_supply,cur_supply");
        tableFactory.createTable(USER_TABLE, "user", "token_name,balance");
    }

    // create repo
    function createRepo(string memory token_name, string memory owner, uint256 total_supply, uint256 cur_supply, string[] memory username, uint256[] memory balance)
        public
        returns (int256)
    {
        if (judgeReposExisted(token_name) == true) {
            return -1001;
        }
        if (total_supply < cur_supply) {
            return -1002;
        }
        Table repo_table = tableFactory.openTable(REPO_TABLE);
        Table user_table = tableFactory.openTable(USER_TABLE);

        Entry repo_entry = repo_table.newEntry();
        repo_entry.set("token_name", token_name);
        repo_entry.set("owner", owner);
        repo_entry.set("total_supply", total_supply);
        repo_entry.set("cur_supply", cur_supply);

        int256 count = repo_table.insert(token_name, repo_entry);
        if (count != 1) {
            return count;
        }

        Entry user_entry = user_table.newEntry();
        uint256 i = 0;
        for (i = 0; i < balance.length; ++i) {
            user_entry = user_table.newEntry();
            user_entry.set("user", username[i]);
            user_entry.set("token_name", token_name);
            user_entry.set("balance", balance[i]);
            count = user_table.insert(username[i], user_entry);
            if (count != 1) {
                return count;
            }
        }
        
        // emit InsertResult(count);
        return count;
    }

    function selectRepoBasicInfo(string memory token_name)
        public
        view
        returns (string memory, string memory, int256, int256)
    {
        Table repo_table = tableFactory.openTable(REPO_TABLE);
        Condition condition = repo_table.newCondition();
        Entries entries = repo_table.select(token_name, condition);
        if (entries.size() == 0) {
            return ("", "", 0, 0);
        }
        Entry entry = entries.get(0);
        return (entry.getString("token_Name"), entry.getString("owner"), entry.getInt("total_supply"), entry.getInt("cur_supply"));
    }

    // select user all balance
    function selectUserAllBalance(string memory username) 
        public
        view
        returns (string[] memory, string[] memory, uint256[] memory)
    {
        Table user_table = tableFactory.openTable(USER_TABLE);

        Condition condition = user_table.newCondition();

        Entries entries = user_table.select(username, condition);
        string[] memory username_list = new string[](uint256(entries.size()));
        string[] memory token_name_list = new string[](uint256(entries.size()));
        uint256[] memory balance_list = new uint256[](uint256(entries.size()));

        for (int256 i = 0; i < entries.size(); ++i) {
            Entry entry = entries.get(i);

            username_list[uint256(i)] = entry.getString("user");
            token_name_list[uint256(i)] = entry.getString("token_name");
            balance_list[uint256(i)] = entry.getUInt("balance");
        }

        return (username_list, token_name_list, balance_list);
    } 

    // select single user balance
    function selectUserBalance(string memory username, string memory token_name) 
        public
        view
        returns (string[] memory, string[] memory, uint256[] memory)
    {
        Table user_table = tableFactory.openTable(USER_TABLE);

        Condition condition = user_table.newCondition();
        condition.EQ("token_name", token_name);

        Entries entries = user_table.select(username, condition);
        string[] memory username_list = new string[](uint256(entries.size()));
        string[] memory token_name_list = new string[](uint256(entries.size()));
        uint256[] memory balance_list = new uint256[](uint256(entries.size()));

        for (int256 i = 0; i < entries.size(); ++i) {
            Entry entry = entries.get(i);

            username_list[uint256(i)] = entry.getString("user");
            token_name_list[uint256(i)] = entry.getString("token_name");
            balance_list[uint256(i)] = entry.getUInt("balance");
        }

        return (username_list, token_name_list, balance_list);
    } 

    // // add user balance
    // function addUserBalance(string memory username, string memory token_name) 
    //     public
    //     view
    //     returns (string[] memory, string[] memory, uint256[] memory)
    // {
    //     Table user_table = tableFactory.openTable(USER_TABLE);

    //     Condition condition = user_table.newCondition();
    //     condition.EQ("token_name", token_name);

    //     Entries entries = user_table.select(username, condition);
    //     string[] memory username_list = new string[](uint256(entries.size()));
    //     string[] memory token_name_list = new string[](uint256(entries.size()));
    //     uint256[] memory balance_list = new uint256[](uint256(entries.size()));

    //     for (int256 i = 0; i < entries.size(); ++i) {
    //         Entry entry = entries.get(i);

    //         username_list[uint256(i)] = entry.getString("user");
    //         token_name_list[uint256(i)] = entry.getString("token_name");
    //         balance_list[uint256(i)] = entry.getUInt("balance");
    //     }

    //     return (username_list, token_name_list, balance_list);
    // } 

    function judgeReposExisted(string memory token_name) private returns (bool) {
        Table repo_table = tableFactory.openTable(REPO_TABLE);
        Condition condition = repo_table.newCondition();
        Entries entries = repo_table.select(token_name, condition);
        if (entries.size() > 0) {
            return true;
        }
        return false;
    }
}