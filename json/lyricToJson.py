
with open('./in.txt', 'r') as inFIle:
    file_data = inFIle.readlines()

    with open('./out.txt', 'w') as out:
        for line in file_data:
            out.write(line[:-1]);
            out.write('\\n');


