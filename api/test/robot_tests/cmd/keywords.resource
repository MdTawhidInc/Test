*** Settings ***
Documentation       Keywords specific to cmd test suite

Library             RequestsLibrary
Library             OperatingSystem
Library             Collections
Resource            ../keywords.resource
Resource            ../variables.resource


*** Keywords ***
Suite Setup
    Basic Suite Setup
    Init SC build environment
    Delete Wallet Plugin
    Install Wallet Plugin

Delete Wallet Plugin
    Log To Console    Deleting Massa Wallet Plugin
    ${pluginId}=    Get Plugin ID From Author and Name    massalabs    wallet
    IF    "${pluginId}" != "${EMPTY}"
        Log To Console    Deleting plugin with ID ${pluginId}
        ${response}=    DELETE    ${API_URL}/plugin-manager/${pluginId}
        Sleep    1 seconds    # Wait for the plugin to be deleted
    END

Install Wallet Plugin
    Log To Console    Installing Massa Wallet Plugin
    ${source}=    Set Variable
    ...    https://github.com/massalabs/thyra-plugin-wallet/releases/download/v0.0.6/thyra-plugin-wallet_${OS}-${ARCH}.zip
    ${response}=    POST
    ...    ${API_URL}/plugin-manager
    ...    params=source=${source}
    ...    expected_status=${STATUS_NO_CONTENT}
    Sleep    1 seconds    # Wait for the plugin to be registered

Init SC build environment
    Log To Console    Initializing SC build environment
    Log To Console    Installing SC dependencies
    Run    cd testSC; npm install
    Log To Console    Building SC
    Run    cd testSC; npm run build

Get SC address
    [Documentation]    Get the address of the SC from a string containing the address after ':' (colon)
    [Arguments]    ${string}
    ${words}=    Split String    ${string}    :
    ${sc_address}=    Get From List    ${words}    1
    RETURN    ${sc_address}