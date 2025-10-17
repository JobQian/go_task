// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    mapping ( address => uint256 ) public votes;
    mapping(address => uint256) private voteVersion;
    uint256 private currentVersion;

    function vote(address adr,uint256 amount) external {
        if (voteVersion[adr] < currentVersion) {
            // 这个地址上次的数据已被逻辑清空
            votes[adr] = amount;
            voteVersion[adr] = currentVersion;
        } else {
            votes[adr] += amount;
        }
    }

    function getVotes(address adr) public view returns (uint256) {
        if (voteVersion[adr] < currentVersion) {
            return 0;
        }
        return votes[adr];
    }

    function resetVotes() internal {
        currentVersion += 1;
    }
}