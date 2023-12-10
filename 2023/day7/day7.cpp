#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>
#include <fstream>
#include <sstream>
#include <algorithm>

#define HAND_SIZE 5

std::string INPUT_PATH = "input.dat";

struct Hand
{
    std::string cards;
    unsigned int bid;
    int value;
};

enum Rank
{
    FIVE_OF_A_KIND = 0,
    FOUR_OF_A_KIND = 1,
    FULL_HOUSE = 2,
    THREE_OF_A_KIND = 3,
    TWO_PAIR = 4,
    ONE_PAIR = 5,
    HIGH_CARD = 6
};

std::vector<Hand> read_hands(const std::string &path)
{
    std::vector<Hand> hands;

    std::ifstream inputFile(path);
    if (!inputFile.is_open())
    {
        std::cerr << "Unable to open file." << std::endl;
        return hands; // Return empty vector if file not opened
    }

    std::string line;
    while (std::getline(inputFile, line))
    {
        std::istringstream iss(line);
        std::string cards;
        unsigned int bid;

        if (iss >> cards >> bid)
        {
            Hand hand;
            hand.cards = cards;
            hand.bid = bid;
            hands.push_back(hand);
        }
        else
        {
            std::cerr << "Invalid input format: " << line << std::endl;
        }
    }

    inputFile.close();
    return hands;
}

int get_value(std::string &hand)
{
    if (hand.size() != HAND_SIZE)
    {
        return -1;
    }

    std::unordered_map<char, int> card_count;

    for (char card : hand)
    {
        card_count[card]++;
    }

    int pair_count = 0;
    for (const auto &pair : card_count)
    {
        if (pair.second == 5)
        {
            return FIVE_OF_A_KIND;
        }
        if (pair.second == 4)
        {
            return FOUR_OF_A_KIND;
        }
        if (pair.second == 3)
        {
            for (const auto &pair2 : card_count)
            {
                if (pair2.second == 2 && pair2.first != pair.first)
                {
                    return FULL_HOUSE;
                }
            }
            return THREE_OF_A_KIND;
        }
        if (pair.second == 2)
        {
            pair_count++;
        }
    }

    if (pair_count == 2)
    {
        return TWO_PAIR;
    }
    if (pair_count == 1)
    {
        return ONE_PAIR;
    }
    return HIGH_CARD;
}

void rank(std::vector<Hand> &hands)
{
    std::sort(hands.begin(), hands.end(), [](const Hand &a, const Hand &b)
              {
                  if (a.value != b.value)
                  {
                      return a.value > b.value;
                  }
                  else
                  {
                      std::string valid_cards = "AKQJT98765432";
                      for (int i = 0; i < HAND_SIZE; i++)
                      {
                          size_t index1 = valid_cards.find(a.cards[i]);
                          size_t index2 = valid_cards.find(b.cards[i]);

                          if (index1 != index2)
                          {
                              return index1 > index2;
                          }
                      }
                  }
                  return false; });
}

void part1(std::vector<Hand> &hands)
{
    int winnings = 0;
    for (Hand &hand : hands)
    {
        hand.value = get_value(hand.cards);
    }
    rank(hands);
    for (int i = 0; i < hands.size(); i++)
    {
        winnings += hands[i].bid * (i + 1);
    }
    std::cout << "Part 1: " << winnings << std::endl;
}

int main()
{
    std::vector<Hand> hands = read_hands(INPUT_PATH);
    part1(hands);
}
