import random


def main():
    board_size = 0
    while(board_size == 0):
        board_size = int(input("Please add Board size: "))
        if(board_size % 5 != 0):
            print("Board size must be multiple of 5.\n")
            board_size = 0
    cardsQty = int(input("Please add number of Cards: "))
    limit = int(board_size / 5)
    for card in range(cardsQty):
        card_board = {'B': [], 'I': [], 'N': [], 'G': [], 'O': []}
        index = 0
        cardNumber = 1000
        for key in card_board:
            numbers_pool = list(range(index * limit + 1, limit * (index + 1) + 1))
            numbers = []
            for i in range(5):
                number = random.choice(numbers_pool)
                numbers_pool.remove(number)
                if key == 'N':
                    if i == 2:
                        number = cardNumber + card
                numbers.append(number)
            card_board[key] = numbers
            index += 1
        filename = "cartones.csv"
        with open(filename, 'a') as csv_file:
            for data in card_board.keys():
                csv_file.write("%s,%s\n" % (data, card_board[data]))


main()
