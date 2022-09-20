#!/usr/bin/env python3
import argparse
import psycopg2
import logging
import subprocess
import sys
import shlex

COCKROACH_EXE = "/home/jennifer/go/src/github.com/cockroachdb/cockroach" \
                "/cockroach"


def call(cmd, stdout=subprocess.PIPE, stderr=subprocess.STDOUT):
  """
    Calls a command in the shell.

    :param cmd: (str)
    :param stdout: set by default to subprocess.PIPE (which is standard stream)
    :param stderr: set by default subprocess.STDOUT (combines with stdout)
    :return: if successful, stdout stream of command.
    """
  print(cmd)
  p = subprocess.run(
    cmd, stdout=stdout, stderr=stderr, shell=True, check=True,
    universal_newlines=True
  )
  return p.stdout


def start_tpcc(server_num, maxservers):
  """ Start a single tpcc server"""

  joins = ",".join(
    ["localhost:{}".format(26257 + i) for i in range(maxservers)]
  )
  options = ["--insecure", "--store=tpcc-local{}".format(server_num + 1),
             "--listen-addr=localhost:{}".format(26257 + server_num),
             "--http-addr=localhost:{}".format(8080 + server_num),
             "--join={}".format(
               joins
             ), "--background", ]

  cmd = "{} start {}".format(
    COCKROACH_EXE, " ".join(options)
  )
  return subprocess.Popen(shlex.split(cmd))


def init_cluster():
  options = ["--insecure", "--host=localhost:26257"]

  cmd = "{} init {}".format(COCKROACH_EXE, " ".join(options))
  _ = call(cmd)


def init_tpcc(warehouses):
  cmd = "{} workload init tpcc --warehouses={} " \
        "'postgresql://root@localhost:26257?sslmode=disable'".format(
    COCKROACH_EXE, warehouses
  )

  _ = call(cmd)


def query_table_num_from_names(names, host="localhost"):
  db_url = "postgresql://root@{}:26257?sslmode=disable".format(host)

  conn = psycopg2.connect(db_url, database="tpcc")

  mapping = {}
  with conn.cursor() as cur:
    for table_name in names:

      query = "SELECT '\"{}\"'::regclass::oid;".format(table_name)
      print(query)

      cur.execute(query)
      logging.debug("status message %s", cur.statusmessage)

      rows = cur.fetchall()
      if len(rows) > 1:
        print("fetchall should only have one row")
        sys.exit(-1)

      mapping[table_name] = rows[0][0]

    conn.commit()

  return mapping


def main():
  parser = argparse.ArgumentParser(description='Process some integers.')
  parser.add_argument(
    "--turnoncicada", action="store_true", default=False,
    help="did you turn on cicada yet?"
  )
  parser.add_argument(
    "--donotrestartserver", dest="restartserver", action="store_false",
    default=True
  )
  parser.add_argument(
    "--numservers", type=int, default=1, help="numberofservers"
  )
  parser.add_argument(
    "--donotinitcluster", dest="initcluster", default=True,
    action="store_false", help="don't init cluster"
  )
  parser.add_argument(
    "--donotinittpcc", dest="inittpcc", action="store_false", default=True,
    help="do not init tpcc"
  )
  parser.add_argument(
    "--warehouses", type=int, default=2, help="num warehouses"
  )
  parser.add_argument(
    "--maptablenums", default=False, action="store_true",
    help="map table numbers"
  )
  args = parser.parse_args()
  if not args.turnoncicada:
    print("TURN ON CICADA")
    return -1

  if args.restartserver:
    total_servers = args.numservers
    processes = []
    for i in range(total_servers):
      processes.append(start_tpcc(i, total_servers))

    for p in processes:
      p.wait()

  if args.initcluster:
    init_cluster()

  if args.inittpcc:
    init_tpcc(args.warehouses)

  if args.maptablenums:
    tableNames = ["warehouse", "stock", "item", "history", "new_order",
                  "order_line", "district", "customer", "order"]
    mapping = query_table_num_from_names(tableNames)
    for name, num in mapping.items():
      print(name, num)

    print("RUN THE POPULATION GO SCRIPT TOO!")

  return 0


if __name__ == "__main__":
  sys.exit(main())
