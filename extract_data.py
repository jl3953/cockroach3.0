#!/usr/bin/env python3
import sys
import argparse


def extract_line_order(line):
  """Sample line:
  W230408 23:46:13.009232 3961 kv/txn.go:2268  jenndebug txn.Id 371425202687
  , AddInsertHotkeys /Table/57/1/0/6/-1877/0, [193 137 136 142 134 248 171 136]
  , [[21 0 0 0 18 0 0 0 193 137 136 142 134 248 171 136 0 23 84 26 98 231 24 1
   151 9 95 47 180 96 10 67 164 35 24 202 131 202 187 8 0 19 20 19 18 19 2]]"""

  comma_sep = line.split(",")

  # table
  table_string = comma_sep[
    1].strip()  # AddInsertHotkeys /Table/57/1/0/6/-1877/0
  table = table_string.split(
    " ")  # [" AddInsertHotkeys", "/Table/57/1/0/6/-1877/0"]
  table_components = table[1].split(
    "/")  # ["", "Table", "57", "1", "0", "6", "-1877", "0"]
  index = int(table_components[3])  # 1
  pkCols = [int(i) for i in
            table_components[4:len(table_components) - 1]]  # [0, 6, -1877]

  # value
  data_string = comma_sep[-1].strip(" \n[]")

  return index, pkCols, data_string


def extract_line_warehouse(line):
  comma_sep = line.split(",")

  # table
  table_string = comma_sep[1].strip(" \n")
  table = table_string.split(" ")
  table_components = table[1].strip(" ").split("/")
  index = int(table_components[3])
  pkCols = [int(i) for i in table_components[4:len(table_components) - 1]]

  # value
  data_string = comma_sep[2].strip("val []\n")
  data_string_length = len(data_string.strip(" \n").split(" "))

  byteData = " ".join([str(i) for i in [data_string_length, 0, 0, 0,
                                        14, 0, 0, 0,
                                        189, 137, 136 + pkCols[0], 136,
                                        0, 23, 82, 151, 8, 105, 47, 38, 122,
                                        9]])

  print(index, pkCols, byteData + " " + data_string)
  return index, pkCols, byteData + " " + data_string


def extract_line(table, line):
  if table == "warehouse":
    return extract_line_warehouse(line)
  elif table == "order":
    return extract_line_order(line)
  else:
    print("nope")
    return ""


def main():
  parser = argparse.ArgumentParser(description='Process some integers.')
  parser.add_argument(
    "datafile", type=str, help="data file"
  )
  parser.add_argument(
    "outfile", type=str, help="output file"
  )
  parser.add_argument(
    "table", type=str, help="table being extracted, i.e. warehouse, order"
  )

  args = parser.parse_args()

  all = []
  with open(args.datafile, "r") as infile:
    for line in infile:
      index, pkCols, byteData = extract_line(args.table, line)
      all.append((index, pkCols, byteData))

  if args.table == "order":
    all = sorted(all, key=lambda x: (x[0], x[1][0], x[1][1], x[1][2]))
  elif args.table == "warehouse":
    all = sorted(all, key=lambda x: (x[0], x[1][0]))
  elif args.table == "district":
    all = sorted(all, key=lambda x: (x[0], x[1][0], x[1][1]))

  print(all)
  with open(args.outfile, "w") as outfile:
    for index, pkCols, byteData in all:
      outfile.write("{0}, {1}, {2}: {3}\n".format(args.table,
                                                  str(index),
                                                  ", ".join(
                                                    [str(i) for i in pkCols]),
                                                  byteData))

  return 0


if __name__ == "__main__":
  sys.exit(main())
