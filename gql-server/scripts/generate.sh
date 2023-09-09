NO_FORMAT="\033[0m"
F_BOLD="\033[1m"
F_DIM="\033[2m"
C_SLATEBLUE1="\033[38;5;99m"
C_LIME="\033[38;5;10m"
C_MEDIUMORCHID1="\033[38;5;171m"

echo -e "${F_BOLD}${C_SLATEBLUE1}⚛ Generating GraphQL in /graph ...${NO_FORMAT}"
echo -e -n "${F_DIM}${C_MEDIUMORCHID1}"
go generate -x ./...
echo -e -n "${NO_FORMAT}"
echo -e "${F_BOLD}${C_LIME}✅ Done!${NO_FORMAT}"
