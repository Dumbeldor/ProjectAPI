#! /usr/bin/env python
# coding: utf-8

import sys
import json

def main():
    if len(sys.argv) < 2:
        print("No destination file given")
        sys.exit(1)

    if len(sys.argv) < 3:
        print("At least one file should be provided.")
        sys.exit(1)

    swagger_result = {}
    files_to_process = sys.argv[2:]
    print("The following files will be merged: {}".format(files_to_process))

    for fn in files_to_process:
        with open(fn) as json_file:
            print("Processing {}".format(fn))

            data = json.load(json_file)
            if len(swagger_result.keys()) == 0:
                swagger_result = data
            else:
                if "paths" in data.keys():
                    for k in data["paths"].keys():
                        swagger_result["paths"][k] = data["paths"][k]
                if "responses" in data.keys():
                    for k in data["responses"].keys():
                        swagger_result["responses"][k] = data["responses"][k]
                if "definitions" in data.keys():
                    for k in data["definitions"].keys():
                        swagger_result["definitions"][k] = data["definitions"][k]

    dest_file = sys.argv[1]
    with open(dest_file, 'w') as outfile:
        print("Writing result to: {}".format(dest_file))
        json.dump(swagger_result, outfile)
if __name__ == "__main__":
    main()
