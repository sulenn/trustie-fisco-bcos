pragma solidity>=0.4.24 <=0.7.0;
pragma experimental ABIEncoderV2;

import "./Table.sol";

contract OpenSource {
    TableFactory tableFactory;
    string constant REPO_TABLE = "repo";
    string constant USER_TABLE = "user";
    string constant COMMIT_TABLE = "commit";
    string constant PUSH_TABLE = "push";
    string constant PULL_REQUEST_TABLE = "pull_request";
    string constant PULL_REQUEST_COMMENT_TABLE = "pull_request_comment";
    string constant ISSUE_TABLE = "issue";
    string constant ISSUE_COMMENT_TABLE = "issue_comment";

    constructor() public {
        tableFactory = TableFactory(0x1001); //The fixed address is 0x1001 for TableFactory
        // the parameters of createTable are tableName,keyField,"vlaueFiled1,vlaueFiled2,vlaueFiled3,..."
        tableFactory.createTable(REPO_TABLE, "token_name", "owner,total_supply,cur_supply");
        tableFactory.createTable(USER_TABLE, "user", "token_name,balance");
        tableFactory.createTable(COMMIT_TABLE, "commit_hash", "content");
        tableFactory.createTable(PUSH_TABLE, "push_id", "content");
        tableFactory.createTable(PULL_REQUEST_TABLE, "pull_request_id", "content");
        tableFactory.createTable(PULL_REQUEST_COMMENT_TABLE, "pull_request_comment_id", "content");
        tableFactory.createTable(ISSUE_TABLE, "issue_id", "content");
        tableFactory.createTable(ISSUE_COMMENT_TABLE, "issue_comment_id", "content");
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

    // add user balance
    function addUserBalance(string memory username, string memory token_name, uint256 amount) 
        public
        returns (int256)
    {
        // 获取项目总Token数，并判断用户增加的 Token 是否导致总 Token 溢出
        Table repo_table = tableFactory.openTable(REPO_TABLE);
        Condition condition = repo_table.newCondition();
        Entries entries = repo_table.select(token_name, condition);
        if (entries.size() == 0) {
            return -1003;
        }
        Entry entry = entries.get(0);
        uint256 cur_supply = entry.getUInt("cur_supply");
        cur_supply += amount;
        if (cur_supply > entry.getUInt("total_supply")) {
            return -1004;
        }

        // 获取用户
        Table user_table = tableFactory.openTable(USER_TABLE);
        condition = user_table.newCondition();
        condition.EQ("token_name", token_name);
        entries = user_table.select(username, condition);

        uint256 user_balance = amount;
        int256 count = 0;
        if (entries.size() == 0) {
            // 用户不存在，插入用户信息
            entry = user_table.newEntry();
            entry.set("user", username);
            entry.set("token_name", token_name);
            entry.set("balance", user_balance);
            count = user_table.insert(username, entry);
            if (count != 1) {
                return count;
            }
        } else {
            // 更新用户信息
            entry = entries.get(0);
            Entry updated_user_entry = user_table.newEntry();
            user_balance += entry.getUInt("balance");
            updated_user_entry.set("balance", user_balance);
            condition = user_table.newCondition();
            condition.EQ("token_name", token_name);
            condition.EQ("user", username);
            count = user_table.update(username, updated_user_entry, condition);
             if (count != 1) {
                return count;
            }
        }
        // 更新 repo 信息
        condition = repo_table.newCondition();
        entry = repo_table.newEntry();
        entry.set("cur_supply", cur_supply);
        condition.EQ("token_name",token_name);
        count = repo_table.update(token_name, entry, condition);
        
        return count;
    } 

    // minus user balance
    function minusUserBalance(string memory username, string memory token_name, uint256 amount) 
        public
        returns (int256)
    {
        // 获取项目总Token数，并判断用户减少的 Token 是否导致项目当前 Token 溢出
        Table repo_table = tableFactory.openTable(REPO_TABLE);
        Condition condition = repo_table.newCondition();
        Entries entries = repo_table.select(token_name, condition);
        if (entries.size() == 0) {
            return -1003;
        }
        Entry entry = entries.get(0);
        uint256 cur_supply = entry.getUInt("cur_supply") - amount;

        // 获取用户
        Table user_table = tableFactory.openTable(USER_TABLE);
        condition = user_table.newCondition();
        condition.EQ("token_name", token_name);
        entries = user_table.select(username, condition);

        uint256 user_balance = amount;
        int256 count = 0;
        if (entries.size() == 0) {
            // 用户不存在，插入用户信息
            entry = user_table.newEntry();
            entry.set("user", username);
            entry.set("token_name", token_name);
            entry.set("balance", "0");
            count = user_table.insert(username, entry);
            if (count != 1) {
                return count;
            }
            return -1006;
        } else {
            // 更新用户信息
            entry = entries.get(0);
            Entry updated_user_entry = user_table.newEntry();
            if  (entry.getUInt("balance") < amount) {
                return -1006;
            }
            user_balance = entry.getUInt("balance") - amount;
            updated_user_entry.set("balance", user_balance);
            condition = user_table.newCondition();
            condition.EQ("token_name", token_name);
            condition.EQ("user", username);
            count = user_table.update(username, updated_user_entry, condition);
             if (count != 1) {
                return count;
            }
        }
        // 更新 repo 信息
        condition = repo_table.newCondition();
        entry = repo_table.newEntry();
        entry.set("cur_supply", cur_supply);
        condition.EQ("token_name",token_name);
        count = repo_table.update(token_name, entry, condition);
        
        return count;
    } 

        // minus user balance
    function transferUserBalance(string memory payer, string memory payee, string memory token_name, uint256 amount) 
        public
        returns (int256)
    {
        // 判断项目是否存在
        Table repo_table = tableFactory.openTable(REPO_TABLE);
        Condition condition = repo_table.newCondition();
        Entries entries = repo_table.select(token_name, condition);
        if (entries.size() == 0) {
            return -1003;
        }

        // 获取转账人信息
        Table user_table = tableFactory.openTable(USER_TABLE);
        condition = user_table.newCondition();
        condition.EQ("token_name", token_name);
        entries = user_table.select(payer, condition);

        int256 count = 0;
        Entry entry = user_table.newEntry();
        if (entries.size() == 0) {
            // 转账人不存在，新增转账人信息
            entry = user_table.newEntry();
            entry.set("user", payer);
            entry.set("token_name", token_name);
            entry.set("balance", "0");
            count = user_table.insert(payer, entry);
            if (count != 1) {
                return count;
            }
            return -1006;
        } 
        entry = entries.get(0);
        Entry updated_entry = user_table.newEntry();
        if  (entry.getUInt("balance") < amount) {
            // 转账人存在，但转账人余额不足
            return -1006;
        }
        
        // 更新转账人余额
        updated_entry.set("balance", entry.getUInt("balance")-amount);
        condition = user_table.newCondition();
        condition.EQ("token_name", token_name);
        condition.EQ("user", payer);
        count = user_table.update(payer, updated_entry, condition);
         if (count != 1) {
            return count;
        }

        // 判断收款人是否存在，存在则更新余额，不存在则新增
        condition = user_table.newCondition();
        condition.EQ("token_name", token_name);
        entries = user_table.select(payee, condition);
        if (entries.size() == 0) {
            // 收款人不存在，新增收款人信息
            entry = user_table.newEntry();
            entry.set("user", payee);
            entry.set("token_name", token_name);
            entry.set("balance", amount);
            count = user_table.insert(payee, entry);
            return count;
        } 
        entry = entries.get(0);
        updated_entry = user_table.newEntry();
        // 更新收款人余额
        updated_entry.set("balance", entry.getUInt("balance")+amount);
        condition = user_table.newCondition();
        condition.EQ("token_name", token_name);
        condition.EQ("user", payee);
        count = user_table.update(payee, updated_entry, condition);

        return count;
    } 

    function judgeReposExisted(string memory token_name) private returns (bool) {
        Table repo_table = tableFactory.openTable(REPO_TABLE);
        Condition condition = repo_table.newCondition();
        Entries entries = repo_table.select(token_name, condition);
        if (entries.size() > 0) {
            return true;
        }
        return false;
    }

    // add commit info
    function addCommitData(string memory commit_hash, string memory content) 
        public
        returns (int256)
    {
        Table commit_table = tableFactory.openTable(COMMIT_TABLE);
        Entry commit_entry = commit_table.newEntry();
        commit_entry.set("commit_hash", commit_hash);
        commit_entry.set("content", content);

        int256 count = commit_table.insert(commit_hash, commit_entry);
        return count;
    }

    // select commit info
    function selectCommitInfo(string memory commit_hash) 
        public
        view
        returns (string memory, string memory)
    {
        Table commit_table = tableFactory.openTable(COMMIT_TABLE);
        Condition condition = commit_table.newCondition();
        Entries entries = commit_table.select(commit_hash, condition);
         if (entries.size() == 0) {
            return ("", "");
        }
        Entry entry = entries.get(entries.size()-1);
        return (entry.getString("commit_hash"), entry.getString("content"));
    } 

    // add push info
    function addPushData(string memory push_id,  string memory content) 
        public
        returns (int256)
    {
        Table push_table = tableFactory.openTable(PUSH_TABLE);
        Entry push_entry = push_table.newEntry();
        push_entry.set("push_id", push_id);
        push_entry.set("content", content);

        int256 count = push_table.insert(push_id, push_entry);
        return count;
    }

    // select push info
    function selectPushInfo(string memory push_id) 
        public
        view
        returns (string memory, string memory)
    {
        Table push_table = tableFactory.openTable(PUSH_TABLE);
        Condition condition = push_table.newCondition();
        Entries entries = push_table.select(push_id, condition);
         if (entries.size() == 0) {
            return ("","");
        }
        Entry entry = entries.get(entries.size()-1);
        return (entry.getString("push_id"),  entry.getString("content"));
    } 

    // add pull request info
    function addPullRequestData(string memory pull_request_id, string memory content) 
        public
        returns (int256)
    {
        Table pull_request_table = tableFactory.openTable(PULL_REQUEST_TABLE);
        Entry pull_request_entry = pull_request_table.newEntry();
        pull_request_entry.set("pull_request_id", pull_request_id);
        pull_request_entry.set("content", content);

        int256 count = pull_request_table.insert(pull_request_id, pull_request_entry);
        return count;
    }

    // select pull request latest info
    function selectPullRequestInfo(string memory pull_request_id) 
        public
        view
        returns (string memory, string memory)
    {
        Table pull_request_table = tableFactory.openTable(PULL_REQUEST_TABLE);
        Condition condition = pull_request_table.newCondition();
        Entries entries = pull_request_table.select(pull_request_id, condition);
         if (entries.size() == 0) {
            return ("", "");
        }
        Entry entry = entries.get(entries.size()-1);
        return (entry.getString("pull_request_id"), entry.getString("content"));
    } 

    // select pull request all info
    function selectPullRequestAllInfo(string memory pull_request_id) 
        public
        view
        returns (string[] memory, string[] memory)
    {
        Table pull_request_table = tableFactory.openTable(PULL_REQUEST_TABLE);
        Condition condition = pull_request_table.newCondition();
        Entries entries = pull_request_table.select(pull_request_id, condition);
        string[] memory pull_request_id_list = new string[](uint256(entries.size()));
        string[] memory content_list = new string[](uint256(entries.size()));
         
        for (int256 i = 0; i < entries.size(); ++i) {
            Entry entry = entries.get(i);

            pull_request_id_list[uint256(i)] = entry.getString("pull_request_id");
            content_list[uint256(i)] = entry.getString("content");
        }
         
        return (pull_request_id_list, content_list);
    } 

    // add pull request comment info
    function addPullRequestCommentData(string memory pull_request_comment_id, string memory content) 
        public
        returns (int256)
    {
        Table pull_request_comment_table = tableFactory.openTable(PULL_REQUEST_COMMENT_TABLE);
        Entry pull_request_comment_entry = pull_request_comment_table.newEntry();
        pull_request_comment_entry.set("pull_request_comment_id", pull_request_comment_id);
        pull_request_comment_entry.set("content", content);

        int256 count = pull_request_comment_table.insert(pull_request_comment_id, pull_request_comment_entry);
        return count;
    }

    // select pull request comment latest info
    function selectPullRequestCommentInfo(string memory pull_request_comment_id) 
        public
        view
        returns (string memory, string memory)
    {
        Table pull_request_comment_table = tableFactory.openTable(PULL_REQUEST_COMMENT_TABLE);
        Condition condition = pull_request_comment_table.newCondition();
        Entries entries = pull_request_comment_table.select(pull_request_comment_id, condition);
         if (entries.size() == 0) {
            return ("", "");
        }
        Entry entry = entries.get(entries.size()-1);
        return (entry.getString("pull_request_comment_id"), entry.getString("content"));
    } 

    // select pull request comment all info
    function selectPullRequestCommentAllInfo(string memory pull_request_comment_id) 
        public
        view
        returns (string[] memory, string[] memory)
    {
        Table pull_request_comment_table = tableFactory.openTable(PULL_REQUEST_COMMENT_TABLE);
        Condition condition = pull_request_comment_table.newCondition();
        Entries entries = pull_request_comment_table.select(pull_request_comment_id, condition);
        string[] memory pull_request_id_comment_list = new string[](uint256(entries.size()));
        string[] memory content_list = new string[](uint256(entries.size()));
         
        for (int256 i = 0; i < entries.size(); ++i) {
            Entry entry = entries.get(i);

            pull_request_id_comment_list[uint256(i)] = entry.getString("pull_request_comment_id");
            content_list[uint256(i)] = entry.getString("content");
        }
         
        return (pull_request_id_comment_list, content_list);
    } 

    // add issue info
    function addIssueData(string memory issue_id, string memory content) 
        public
        returns (int256)
    {
        Table issue_table = tableFactory.openTable(ISSUE_TABLE);
        Entry issue_entry = issue_table.newEntry();
        issue_entry.set("issue_id", issue_id);
        issue_entry.set("content", content);

        int256 count = issue_table.insert(issue_id, issue_entry);
        return count;
    }

    // select issue latest info
    function selectIssueInfo(string memory issue_id) 
        public
        view
        returns (string memory, string memory)
    {
        Table issue_table = tableFactory.openTable(ISSUE_TABLE);
        Condition condition = issue_table.newCondition();
        Entries entries = issue_table.select(issue_id, condition);
         if (entries.size() == 0) {
            return ("", "");
        }
        Entry entry = entries.get(entries.size()-1);
        return (entry.getString("issue_id"), entry.getString("content"));
    } 

    // select issue all info
    function selectIssueAllInfo(string memory issue_id) 
        public
        view
        returns (string[] memory, string[] memory)
    {
        Table issue_table = tableFactory.openTable(ISSUE_TABLE);
        Condition condition = issue_table.newCondition();
        Entries entries = issue_table.select(issue_id, condition);
        string[] memory issue_list = new string[](uint256(entries.size()));
        string[] memory content_list = new string[](uint256(entries.size()));
         
        for (int256 i = 0; i < entries.size(); ++i) {
            Entry entry = entries.get(i);

            issue_list[uint256(i)] = entry.getString("issue_id");
            content_list[uint256(i)] = entry.getString("content");
        }
         
        return (issue_list, content_list);
    } 

    // add issue comment info
    function addIssueCommentData(string memory issue_comment_id, string memory content) 
        public
        returns (int256)
    {
        Table issue_comment_table = tableFactory.openTable(ISSUE_COMMENT_TABLE);
        Entry issue_comment_entry = issue_comment_table.newEntry();
        issue_comment_entry.set("issue_comment_id", issue_comment_id);
        issue_comment_entry.set("content", content);

        int256 count = issue_comment_table.insert(issue_comment_id, issue_comment_entry);
        return count;
    }

    // select issue comment latest info
    function selectIssueCommentInfo(string memory issue_comment_id) 
        public
        view
        returns (string memory, string memory)
    {
        Table issue_comment_table = tableFactory.openTable(ISSUE_COMMENT_TABLE);
        Condition condition = issue_comment_table.newCondition();
        Entries entries = issue_comment_table.select(issue_comment_id, condition);
         if (entries.size() == 0) {
            return ("", "");
        }
        Entry entry = entries.get(entries.size()-1);
        return (entry.getString("issue_comment_id"), entry.getString("content"));
    } 

    // select issue comment all info
    function selectIssueCommentAllInfo(string memory issue_comment_id) 
        public
        view
        returns (string[] memory, string[] memory)
    {
        Table issue_comment_table = tableFactory.openTable(ISSUE_COMMENT_TABLE);
        Condition condition = issue_comment_table.newCondition();
        Entries entries = issue_comment_table.select(issue_comment_id, condition);
        string[] memory issue_comment_list = new string[](uint256(entries.size()));
        string[] memory content_list = new string[](uint256(entries.size()));
         
        for (int256 i = 0; i < entries.size(); ++i) {
            Entry entry = entries.get(i);

            issue_comment_list[uint256(i)] = entry.getString("issue_comment_id");
            content_list[uint256(i)] = entry.getString("content");
        }
         
        return (issue_comment_list, content_list);
    } 
}