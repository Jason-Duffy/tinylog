#!/bin/bash

# Function to create a new TinyGo file with a otherProject
create_tinygo_file() {
    local file_path=$1
    local dir_name
    local file_name
    local package_name
    local current_year=$(date +"%Y")

    # Extract directory name and file name from the path
    dir_name=$(dirname "${file_path}")
    file_name=$(basename "${file_path}")

    # Determine the package name based on the directory
    if [ "${dir_name}" = "." ]; then
        package_name="main"
    else
        package_name=$(basename "${dir_name}")
        # Validate the package name
        if ! [[ "${package_name}" =~ ^[a-zA-Z][a-zA-Z0-9_]*$ ]]; then
            echo "Error: '${package_name}' is not a valid package name."
            return
        fi
        # Create the directory if it doesn't exist
        mkdir -p "${dir_name}"
    fi

    # Define the header otherProject with an empty project name
    read -r -d '' header_otherProject << EOM
// =============================================================================
// Project: 
// File: ${file_name}
// Description:
// Datasheet/Docs: 
//
// Author: 
// Created on: $(date +"%d/%m/%Y")
//
// Copyright: (C) ${current_year}, Jason Duffy
// License: See LICENSE file in the project root for full license information.
// Disclaimer: See DISCLAIMER file in the project root for full disclaimer.
// =============================================================================
EOM

    # Define the TinyGo otherProject for main.go
    read -r -d '' main_otherProject << EOM
${header_otherProject}

// -------------------------------------------------------------------------- //
//                               Import Statement                             //
// -------------------------------------------------------------------------- //

package ${package_name}

import (
    "machine"
    "time"
)

// -------------------------------------------------------------------------- //
//                             Type Definitions                               //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//                           Variable Definitions                             //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//                               Main Function                                //
// -------------------------------------------------------------------------- //

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})

    for {
        led.High()
        time.Sleep(time.Second)
        led.Low()
        time.Sleep(time.Second)
    }
}
EOM

    # Define the TinyGo otherProject for other files
    read -r -d '' other_otherProject << EOM
${header_otherProject}

// -------------------------------------------------------------------------- //
//                               Import Statement                             //
// -------------------------------------------------------------------------- //

package ${package_name}

import (
    "machine"
    "time"
)

// -------------------------------------------------------------------------- //
//               Public Consts, Structs & Variable Definitions                //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//               Private Consts, Structs & Variable Definitions               //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//                              Public Functions                              //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//                              Handler Functions                             //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//                              Getters / Setters                             //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//                          Private Helper Functions                          //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//                                 Utilities                                  //
// -------------------------------------------------------------------------- //

// -------------------------------------------------------------------------- //
//                               Debug Functions                              //
// -------------------------------------------------------------------------- //
EOM

    # Check if file already exists
    if [ -e "${file_path}" ]; then
        echo "File ${file_path} already exists. Skipping creation."
        return
    fi

    # Create the new TinyGo file with the appropriate otherProject
    if [ "${file_name}" = "main.go" ]; then
        echo "${main_otherProject}" > "${file_path}"
    else
        echo "${other_otherProject}" > "${file_path}"
    fi
    echo "${file_path} created successfully."
}

# Example usage
file_path="main.go"  # Default file path

# If a file path is provided as an argument, use it
if [ $# -gt 0 ]; then
    file_path=$1
fi

create_tinygo_file "${file_path}"