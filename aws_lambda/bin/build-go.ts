#!/usr/bin/env node
import "source-map-support/register";
import util = require("util");
import generateUid = require("nanoid/generate");

const exec = util.promisify(require("child_process").exec);
const uid: string = generateUid("1234567890abcdefghijklmnopqrstuvwxyz", 13);

/**
 * Set UID
 */
async function setFile(uid: string) {
  await exec(`mkdir -p tmp/cache && echo ${uid} >tmp/cache/uid`);
}
setFile(uid);

/**
 * Build Golang
 */
async function buildGo(pkg: string, uid: string) {
  await exec(`(
    cd lambda/${pkg}
    rm -fr _build
    mkdir -p _build
    GOOS=linux GOARCH=amd64 go build -o _build/${pkg}-${uid}
    zip _build/${pkg}-${uid}.zip _build/${pkg}-${uid}
  )`);
}
buildGo("authorizer", uid);
buildGo("pets", uid);
