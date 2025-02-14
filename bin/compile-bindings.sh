#!/bin/bash

BINDING_DIR=./pkg/bindings
JSON_DIR=./out

function create_binding {
    contract_name=$1

    mkdir -p $BINDING_DIR/${contract_name}

    contract_json_path="${JSON_DIR}/${contract_name}.sol/${contract_name}.json"

    binding_out_dir="${BINDING_DIR}/${contract_name}"
    mkdir -p $binding_out_dir || true

    cat $contract_json_path | jq -r '.abi' > $binding_out_dir/tmp.abi
    cat $contract_json_path | jq -r '.bytecode.object' > $binding_out_dir/tmp.bin

    abigen \
        --bin=$binding_out_dir/tmp.bin \
        --abi=$binding_out_dir/tmp.abi \
        --pkg="${contract_name}" \
        --out=$BINDING_DIR/$contract_name/binding.go \
        > /dev/null 2>&1

    if [[ $? == "1" ]];
    then
        echo "Failed to generate binding for $contract_json_path"
    fi
    rm $binding_out_dir/tmp.abi
    rm $binding_out_dir/tmp.bin
}

# Create or overwrite the README.md file with the initial content
cat << EOF > $BINDING_DIR/README.md
# Go-ABI Bindings for ASP Contracts V1.0

> These bindings were generated with [AbiGen](https://geth.ethereum.org/docs/tools/abigen).

EOF


contracts=$(find src/ -type f -name "*.sol" )
IFS=$'\n'

for contract_name in $contracts; do
    # Append the contract link to README.md
	contract_name=$(basename $contract_name .sol)

    echo "* [$contract_name]($contract_name/)" >> $BINDING_DIR/README.md
	create_binding $contract_name
done
