#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <sstream>
#include <algorithm>
#include <numeric>

#define FILE_IN "input.dat"

struct Block
{
    int start;
    int end;

    int size() const
    {
        return abs(end - start) + 1;
    }
};

struct Record
{
    std::string row;
    std::vector<int> block_sizes;
};

void get_permutations(std::string s, int index, std::vector<std::string> &result)
{
    if (index == s.length())
    {
        result.push_back(s);
        return;
    }

    if (s[index] == '?')
    {
        s[index] = '.';
        get_permutations(s, index + 1, result);
        s[index] = '?';
        get_permutations(s, index + 1, result);
    }
    else
    {
        get_permutations(s, index + 1, result);
    }
}

std::vector<Record> read_records(const std::string &path)
{
    std::vector<Record> records;
    std::ifstream inputFile(path);

    if (!inputFile.is_open())
    {
        std::cerr << "Unable to open file." << std::endl;
    }

    std::string line;
    while (std::getline(inputFile, line))
    {
        Record record;
        std::istringstream iss(line);
        std::string segment;

        iss >> record.row;

        while (std::getline(iss, segment, ' '))
        {
            if (segment.find(',') != std::string::npos)
            {
                std::istringstream block_sizes_stream(segment);
                std::string token;
                Block block;
                while (std::getline(block_sizes_stream, token, ','))
                {
                    record.block_sizes.push_back(std::stoi(token));
                }
            }
        }

        records.push_back(record);
    }

    inputFile.close();
    return records;
}

bool fits(const std::string &row, Block &block)
{
    if (block.end >= row.size() || block.start < 0)
    {
        return false;
    }
    auto found = std::find(row.begin() + block.start, row.begin() + block.end + 1, '.');
    return found == row.begin() + block.end + 1;
}

bool arrangement(const std::string &row, const std::vector<int> &block_sizes, int index = 0, int block_count = 0)
{
    if (block_count == block_sizes.size())
    {
        return true;
    }

    if (index >= row.size())
    {
        return false;
    }

    for (int i = index; i <= row.size(); ++i)
    {
        Block block;
        block.start = i;
        block.end = i + block_sizes[block_count] - 1;

        if (fits(row, block))
        {
            if (arrangement(row, block_sizes, block.end + 2, block_count + 1))
            {
                return true;
            }
        }
    }

    return false;
}

int arrangements(Record &record)
{
    std::vector<std::string> permutations;
    get_permutations(record.row, 0, permutations);

    int count = 0;
    int sum = std::accumulate(record.block_sizes.begin(), record.block_sizes.end(), 0);

    if (record.block_sizes.size() == 0)
    {
        return -1;
    }

    for (auto &perm : permutations)
    {
        std::replace(perm.begin(), perm.end(), '?', '#');
        int signs = std::count(perm.begin(), perm.end(), '#');
        if (signs != sum)
        {
            continue;
        }
        if (arrangement(perm, record.block_sizes))
        {
            ++count;
        }
    }

    return count;
}

int main()
{
    std::vector<Record> records = read_records(FILE_IN);
    int res = 0;
    for (Record &record : records)
    {
        res += arrangements(record);
    }
    std::cout << "Part 1: " << res << std::endl;
    return 0;
}
