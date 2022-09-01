// GENERATED BY textFileToGoConst
// GitHub:     github.com/logrusorgru/textFileToGoConst
// input file: html\front\website.css
// generated:  Wed Aug 24 15:02:56 CEST 2022

package website

const CSS = `body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #fff;

  background :linear-gradient(90deg,#172230 50.03%,#0c1219 85.82%);
}

a {
  color: white;
}

a:hover {
  color: rgb(164, 164, 164);
}

.massa-logo-banner{
width: 170px;
}

.massa-logo-spinner{
  width: 25px;
}
.spinner-border{
  color: transparent;
}

.loading{
  display: none;
}

.title-wallet{
margin-bottom: 45px;
}
option{
background-color: transparent;
}
select.form-select:focus{
background-color :black;
color :white;

}

.form-control, .form-control:focus{
color :white;
border-radius: 0;
border: none;
border-bottom: 2px solid hsla(0,0%,100%,.2);
background-color: transparent;
box-shadow: none;
}

.form-select {
background-color: transparent;
color: white;
}

h1{
  font-size: 3.5rem;
  color : hsla(0,0%,100%,.2);
}
h2{
color : hsla(0,0%,100%,.2);
}

button {
  border:none;
  padding: 3px 30px;
  border-radius: 100px;
  color: #fff;
}

.primary-button {
  font-size: 1.3rem;
  background-color: #e74e4e;
  border: 3px solid #e74e4e;
}

.alert-danger{
  display: none;
  position: fixed;
  top: 20px; 
  left: 50%;
  transform: translate(-50%, 0);
  margin: auto;
  width: 300px;
  padding: 10px 0;
  background-color: #e74e4e;
  border : none;
  text-align: center;
  color:white;
}

.alert-primary{
display: none;
position: fixed;
top: 20px; 
left: 50%;
transform: translate(-50%, 0);
margin: auto;
width: 600px;
padding: 10px 0;
background-color: hsla(0,0%,100%,.05);
border : none;
text-align: center;
color:white;
}
.clipboard{
cursor: pointer;
margin-left: 10px;
margin-bottom: 4px;
}

.quit-button{
cursor: pointer;
}

#website-deployers-table{
color: white;
}

.table{
color:white;
}

tbody, td, tfoot, th, thead, tr{
border-color:rgb(86, 86, 86);
}

.table>:not(caption)>*>*{
padding: 18px 1.5rem ;
}

.table-striped>tbody>tr:nth-of-type(odd)>*{
background-color: hsla(0,0%,100%,.05);
color:white;
}


/* Popover css */

.popover__wrapper{
  right: 20px;
  top : 70px;
  position: fixed;
  width: 150px;
  text-align: center;
  text-decoration: none;
}

.wallet_button{
  text-decoration: none;

  
}
.popover__title {
  text-align: center;
  font-size: 24px;
  color: white;
  border-radius: 100px;
  padding: 7px 0px;
  background-color: #e74e4e;
}

.popover__content {
  opacity: 0;
  display: none;
  background-color: #e74e4e;
  width:130px;
}
.popover__content:before {
  z-index: -1;
  content: "";
}
.popover__wrapper:hover .popover__content {
  z-index: 10;
  opacity: 1;
  display: inline-block;
}

.wallet-item{
  padding: 2px 3px;
}

.wallet-link{
  text-decoration: none;
  color : white;
}

#wallet-list{
  list-style-type: none;
  padding: 0px;
}

td,th {
  text-align: center;
  vertical-align: middle;
}
.massa-logo-spinner { 
  width: 25px;
  animation: spin 2s infinite linear;
  -webkit-animation: spin 2s infinite linear;
}
@-webkit-keyframes spin {
0%  {-webkit-transform: rotate(0deg);}
100% {-webkit-transform: rotate(360deg);}	
}

/* PASSWORD MODAL */
.modal-content{
  background: linear-gradient(90deg,#172230 50.03%,#0c1219 85.82%);
}

.close{
  background: transparent;
}
.modal-footer{
  border-top: none;
}`