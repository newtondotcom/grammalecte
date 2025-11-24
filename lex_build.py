#!python3

"""
Lexicon builder
"""

import argparse
import os

import graphspell.dawg as fsa


def build (spfSrc, sLangCode, sLangName, sfDict, bJavaScript=False, sDicName="", sDescription="", sFilter="", cStemmingMethod="S", nCompressMethod=1):
    "transform a text lexicon as a binary indexable dictionary"
    oDAWG = fsa.DAWG(spfSrc, cStemmingMethod, sLangCode, sLangName, sDicName, sDescription, sFilter)
    os.makedirs("graphspell/_dictionaries", exist_ok=True)
    oDAWG.writeAsJSObject("graphspell/_dictionaries/" + sfDict + ".json")
    if bJavaScript:
        os.makedirs("graphspell-js/_dictionaries", exist_ok=True)
        oDAWG.writeAsJSObject("graphspell-js/_dictionaries/" + sfDict + ".json")


def main ():
    "parse args from CLI"
    xParser = argparse.ArgumentParser()
    xParser.add_argument("src_lexicon", type=str, help="path and file name of the source lexicon")
    xParser.add_argument("lang_code", type=str, help="language code")
    xParser.add_argument("lang_name", type=str, help="language name")
    xParser.add_argument("dic_filename", type=str, help="dictionary file name (without extension)")
    xParser.add_argument("-js", "--json", help="Build dictionary in JSON", action="store_true")
    xArgs = xParser.parse_args()
    build(xArgs.src_lexicon, xArgs.lang_code, xArgs.lang_name, xArgs.dic_filename, xArgs.json)


if __name__ == '__main__':
    main()
