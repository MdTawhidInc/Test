// GENERATED BY textFileToGoConst
// GitHub:     github.com/logrusorgru/textFileToGoConst
// input file: html\front\event-manager.js
// generated:  Wed Sep 21 23:57:21 CEST 2022

package front

const EventListener = "//TODO \tEvents must be managed by Operation Id\r\n// The retry mechanism must be done by mimic the same operation of the operation ID related to the event in\r\nclass EventManager {\r\n\tasync subscribe(str, address, callback) {\r\n\t\taxios({\r\n\t\t\turl: `/thyra/events/${str}/${address}`,\r\n\t\t\tmethod: 'GET',\r\n\t\t})\r\n\t\t\t.then((resp) => {\r\n\t\t\t\tcallback(resp);\r\n\t\t\t})\r\n\t\t\t.catch((e) => {\r\n\t\t\t\t// TODO Implement retry mechanism here\r\n\t\t\t\tconsole.error(e);\r\n\t\t\t});\r\n\t}\r\n}\r\n"