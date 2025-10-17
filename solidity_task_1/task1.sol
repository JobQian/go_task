// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//···································01
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


contract task01{
    bytes array = new bytes{1,2,3,4,5,6,34,1,24,53,1,1,243,35,46,57,7,87,,45,77,89};
    //···································02
    //只在英文环境下有效
    function test(string calldata str) public pure returns (string memory){
        string memory result;

        bytes memory bytestr = bytes(str);
        bytes memory bytestrnew = new bytes(bytestr.length);

        for(uint256 i = 0;i<bytestr.length;i++){
                bytestrnew[bytestr.length-1-i] = bytestr[i];
        }

        result = string(bytestrnew);
        return result;
    }
    //支持英文，中文和emoji
    function reverse(string memory str) public pure returns (string memory) {
        bytes memory b = bytes(str);
        uint len = b.length;

        // 临时存放每个 UTF-8 字符
        string[] memory chars = new string[](len);
        uint count = 0;

        // 解析 UTF-8
        for (uint i = 0; i < len; i++) {
            uint8 first = uint8(b[i]);
            uint charLen;
            if ((first >> 7) == 0) {
                charLen = 1;              // 0xxxxxxx
            } else if ((first >> 5) == 0x06) {
                charLen = 2;              // 110xxxxx
            } else if ((first >> 4) == 0x0E) {
                charLen = 3;              // 1110xxxx
            } else if ((first >> 3) == 0x1E) {
                charLen = 4;              // 11110xxx
            } else {
                // 非合法 UTF-8 首字节，作为 1 字节处理以避免 revert（可根据需要改为 revert）
                charLen = 1;
            }

            bytes memory ch = new bytes(charLen);
            for (uint j = 0; j < charLen; j++) {
                ch[j] = b[i + j];
            }
            chars[count] = string(ch);
            count++;
            i += (charLen - 1);
        }

        // 按字符反转拼接
        string memory result = "";
        for (uint i = count; i > 0; i--) {
            result = string(abi.encodePacked(result, chars[i - 1]));
        }

        return result;
    }
    //···································03
    function inttoroman(int num) public pure returns(string memory) {
        return "";
    }
    //···································04
    function romantoint(string memory momanstr) public pure returns(int256){
        return 0;
    }
    //···································05
    function mergeSortedArray(bytes memory first,bytes memory second) public pure returns(bytes memory){
                return abi.encodePacked(first, second);
    }
    function mergeSortedArray(uint[] memory first,uint[] memory second) public pure returns(uint[] memory){
        uint[] memory result = new uint[](a.length + b.length);
        uint k = 0;

        for (uint i = 0; i < a.length; i++) {
            result[k++] = a[i];
        }

        for (uint i = 0; i < b.length; i++) {
            result[k++] = b[i];
        }

        return result;
    }
    function mergeSortedArray(string memory first,string memory second) public pure returns(string memory){
            return string(abi.encodePacked(first, second));
    }
    //···································06
    function binarySearch(int taget) public view returns (int){

    }
}