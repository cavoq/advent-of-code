#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <sstream>
#include <queue>

#define FILE_IN "input.dat"

struct Point
{
    int x, y;
};

std::vector<Point> get_galaxies(const std::vector<std::string> &grid)
{
    std::vector<Point> points;
    for (int i = 0; i < grid.size(); ++i)
    {
        for (int j = 0; j < grid[i].size(); ++j)
        {
            if (grid[i][j] == '#')
            {
                Point point;
                point.x = i;
                point.y = j;
                points.push_back(point);
            }
        }
    }
    return points;
}

std::vector<std::pair<Point, Point>> get_unique_galaxy_pairs(const std::vector<Point> &galaxies)
{
    std::vector<std::pair<Point, Point>> pairs;

    for (int i = 0; i < galaxies.size(); ++i)
    {
        for (int j = i + 1; j < galaxies.size(); ++j)
        {
            if (galaxies[i].x != galaxies[j].x || galaxies[i].y != galaxies[j].y)
            {
                pairs.push_back({galaxies[i], galaxies[j]});
            }
        }
    }

    return pairs;
}

bool expand_grid(std::vector<std::string> &grid)
{
    if (grid.size() == 0)
    {
        std::cerr << "Empty grid." << std::endl;
        return false;
    }

    int rows = grid.size();
    int cols = grid[0].size();

    for (int i = 0; i < rows; ++i)
    {
        if (grid[i].find('#') == std::string::npos)
        {
            grid.insert(grid.begin() + i, std::string(cols, '.'));
            ++rows;
            ++i;
        }
    }

    for (int j = 0; j < cols; ++j)
    {
        bool expand_column = true;
        for (int i = 0; i < rows; ++i)
        {
            if (grid[i][j] == '#')
            {
                expand_column = false;
                break;
            }
        }
        if (expand_column)
        {
            for (int i = 0; i < rows; ++i)
            {
                grid[i].insert(grid[i].begin() + j, '.');
            }
            ++cols;
            ++j;
        }
    }
    return true;
}

std::pair<std::vector<int>, std::vector<int>> get_expanded_rows_cols(const std::vector<std::string> &grid)
{
    int rows = grid.size();
    int cols = grid[0].size();

    std::vector<int> expanded_rows(rows, 0);
    std::vector<int> expanded_cols(cols, 0);

    for (int i = rows - 1; i >= 0; --i)
    {
        if (grid[i].find('#') == std::string::npos)
        {
            expanded_rows[i] = 1;
        }
    }

    for (int j = cols - 1; j >= 0; --j)
    {
        bool expand_column = true;
        for (int i = 0; i < rows; ++i)
        {
            if (grid[i][j] == '#')
            {
                expand_column = false;
                break;
            }
        }
        if (expand_column)
        {
            expanded_cols[j] = 1;
            ++cols;
        }
    }

    return {expanded_rows, expanded_cols};
}

int shortest_path(const Point &start, const Point &end)
{
    return abs(start.x - end.x) + abs(start.y - end.y);
}

int shortest_path(const Point &start, const Point &end,
                  const std::vector<int> &expanded_rows,
                  const std::vector<int> &expanded_cols,
                  int expansion_factor)
{
    int shortest_path = 0;

    int range_x = abs(start.x - end.x);
    int range_y = abs(start.y - end.y);

    int start_x = start.x < end.x ? start.x : end.x;
    int start_y = start.y < end.y ? start.y : end.y;

    for (int i = start_x; i < start_x + range_x; ++i)
    {
        shortest_path += 1;
        if (expanded_rows[i] == 1)
        {   
            int exp_factor = expansion_factor == 1 ? 2 : expansion_factor;
            shortest_path += exp_factor - 1;
        }
    }

    for (int j = start_y; j < start_y + range_y; ++j)
    {
        shortest_path += 1;
        if (expanded_cols[j] == 1)
        {
            int exp_factor = expansion_factor == 1 ? 2 : expansion_factor;
            shortest_path += exp_factor - 1;
        }
    }

    return shortest_path;
}

std::vector<std::string> read_grid(const std::string &path)
{
    std::vector<std::string> grid;
    std::ifstream inputFile(path);

    if (!inputFile.is_open())
    {
        std::cerr << "Unable to open file." << std::endl;
    }

    std::string line;
    while (std::getline(inputFile, line))
    {
        grid.push_back(line);
    }

    inputFile.close();
    return grid;
}

long long int part1(std::vector<std::string> &grid, int exp_factor = 1)
{
    std::pair<std::vector<int>, std::vector<int>> expanded_rows_cols = get_expanded_rows_cols(grid);
    std::vector<Point> points = get_galaxies(grid);
    std::vector<std::pair<Point, Point>> pairs = get_unique_galaxy_pairs(points);

    long long int sum = 0;
    for (int i = 0; i < pairs.size(); ++i)
    {
        sum += shortest_path(pairs[i].first, pairs[i].second,
                               expanded_rows_cols.first, expanded_rows_cols.second, exp_factor);
    }

    return sum;
}

int main()
{
    std::vector<std::string> grid = read_grid(FILE_IN);
    long long int p1 = part1(grid, 1);
    long long int p2 = part1(grid, 10000000);
    std::cout << "Part 1: " << p1 << std::endl;
    std::cout << "Part 2: " << p2 << std::endl;
}
