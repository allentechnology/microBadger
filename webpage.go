package main

var webpage = `
<html>
    <head>
	<title>MicroBadger</title>
	<style>
	 #login-area {
	     width: 500px;
	     float: left;
	 }
	 #notification-area {
	     margin-left: 500px;
	 }
	 table {
	     width: 100%;
	 }

	 .acidjs-css3-treeview{
	     overflow-y:scroll;
	     wite-space: nowrap;
	     height:300px;
	     border: 1px solid gray;
	     border-collapse: collapse;
	 }



	 /*
	  * Imageless CSS3 Treeview with Checkbox Support
	  * @namespace window.AcidJs
	  * @class CSS3Treeview
	  * @version 3.0
	  * @author Martin Ivanov
	  * @url developer website: http://wemakesites.net/
	  * @url developer twitter: https://twitter.com/#!/wemakesitesnet
	  * @url developer blog http://acidmartin.wordpress.com/
	  **/
	 
	 /*
	  * Do you like this solution? Please, donate:
	  * https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=QFUHPWJB2JDBS
	  **/
	 
	 .acidjs-css3-treeview,
	 .acidjs-css3-treeview *
	 {
	     padding: 0;
	     margin: 0;
	     list-style: none;
	 }
	 
	 .acidjs-css3-treeview label[for]::before,
	 .acidjs-css3-treeview label span::before
	 {
	     content: "\25b6";
	     display: inline-block;
	     margin: 2px 0 0;
	     width: 13px;
	     height: 13px;
	     vertical-align: top;
	     text-align: center;
	     color: #e74c3c;
	     font-size: 8px;
	     line-height: 13px;
	 }
	 
	 .acidjs-css3-treeview li ul
	 {
	     margin: 0 0 0 22px;
	 }
	 
	 .acidjs-css3-treeview *
	 {
	     vertical-align: middle;
	 }
	 
	 .acidjs-css3-treeview
	 {
	     font: normal 11px/16px "Segoe UI", Arial, Sans-serif;
	 }
	 
	 .acidjs-css3-treeview li
	 {
	     -webkit-user-select: none;
	     -moz-user-select: none;
	     user-select: none;
	 }
	 
	 .acidjs-css3-treeview input[type="checkbox"]
	 {
	     display: none;
	 }
	 
	 .acidjs-css3-treeview label
	 {
	     cursor: pointer;
	 }
	 
	 .acidjs-css3-treeview label[for]::before
	 {
	     -webkit-transform: translatex(-24px);
	     -moz-transform: translatex(-24px);
	     -ms-transform: translatex(-24px);
	     -o-transform: translatex(-24px);
	     transform: translatex(-24px);
	 }
	 
	 .acidjs-css3-treeview label span::before
	 {
	     -webkit-transform: translatex(16px);
	     -moz-transform: translatex(16px);
	     -ms-transform: translatex(16px);
	     -o-transform: translatex(16px);
	     transform: translatex(16px);
	 }
	 
	 .acidjs-css3-treeview input[type="checkbox"][id]:checked ~ label[for]::before
	 {
	     content: "\25bc";
	 }
	 
	 .acidjs-css3-treeview input[type="checkbox"][id]:not(:checked) ~ ul
	 {
	     display: none;
	 }
	 
	 .acidjs-css3-treeview label:not([for])
	 {
	     margin: 0 8px 0 0;
	 }
	 
	 .acidjs-css3-treeview label span::before
	 {
	     content: "";
	     border: solid 1px #1375b3;
	     color: #1375b3;
	     opacity: .50;
	 }
	 
	 .acidjs-css3-treeview label input:checked + span::before
	 {
	     content: "\2714";
	     box-shadow: 0 0 2px rgba(0, 0, 0, .25) inset;
	     opacity: 1;
	 }

	 
	 
	</style>
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
	<script>
	 function subForm (postUrl, formID, message){
	     $.ajax({
		 url:"" + postUrl,
		 type:'post',
		 data:$("#" + formID).serialize(),
		 success:function(){
		     alert("" + message);
		 }
	     });
	 }
	 function addContent(frm) {
	     $.post(
		 "/slotSubmit",
		 $('#' + frm).serialize(),
		 function (data) {
                     result = data;
		 }
             )
              .success(function() {
		  alert("hello");
              })
              .complete(function() { 

              })
              .error(function() {
		  alert('An error has occurred.');
              });      

	     return false;// this stops the form from actually posting
	 }
	 function autoRefresh(id){

	     setTimeout(function(){
		 $('#' + id).load(document.URL +  ' #' + id);
		 autoRefresh(id);
	     }, 10000);
	 }

	 function checkSubBoxes(boxID, classID){
	     $("." + classID).prop("checked",$("#" + boxID).is(":checked"));
	     $("." + classID).change();
	 }

	 
	 /* $(".acidjs-css3-treeview").delegate("label input:checkbox", "change", function() {
	    var
	    checkbox = $(this),
	    nestedList = checkbox.parent().next().next(),
	    selectNestedListCheckbox = nestedList.find("label:not([for]) input:checkbox");
	    
	    if(checkbox.is(":checked")) {
	    return selectNestedListCheckbox.prop("checked", true);
	    }
	    selectNestedListCheckbox.prop("checked", false);
	    });*/

	</script>
    </head>
    <body>
	<h1>MicroBadger</h1>
	<div id="menu">
	    <form>
		<!-- <button type="button">Help</button> -->
	    <button type="submit" id="quit-button" title="Quit microBadger and stop randomizing microbadges" formaction="/quit">Quit</button>
	    </form>
	</div>
	<br />
	<div id="login-area">
	    <form action="/login" method="post" id="login-form">
		Username: 
		<input type="text" name="username" autofocus/>
		<br />
		Password: 
		<input type="password" name="password" id="password-field"/>
		<br />
		<!-- <button type="button" onClick="login()" id="login-button" title="Save the login information locally for future use">Save Login</button> -->
		<button type="button" onClick="login()" id="login-button" title="Log into boardgamegeek.com and start randomizing microbadges">Login</button>
		<script>
		 $(document).ready(function(){
		     $('#password-field').keypress(function(e){
			 if(e.keyCode==13)
			     $('#login-button').click();
		     });
		 });

		 function login(){
		     $.ajax({
			 url:'/login',
			 type:'post',
			 data:$('#login-form').serialize()
		     });
		     /* autoRefresh('slot-1');
			autoRefresh('slot-2');
			autoRefresh('slot-3');
			autoRefresh('slot-4');
			autoRefresh('slot-5');*/
		 }

		</script>
		<!-- <input type="submit" value="Save Login" /> -->
	    </form>
	    <!-- <img src="/slot/1" id="slot-1-display" />
		 <img src="/slot/2" id="slot-2-display" />
		 <img src="/slot/3" id="slot-3-display" />
		 <img src="/slot/4" id="slot-4-display" />
		 <img src="/slot/5" id="slot-5-display" /> -->
	</div>
	<div id="notification-area" >
	    <iframe src="/notification"></iframe>
	</div>

	<div id="slots">
	    <table>
		<form action="/slotSubmit" method="post" id="slot-submit-form" onSubmit="addContent('slot-submit-form')">
		    <tr>
			<th>Slot 1</th>
			<th>Slot 2</th>
			<th>Slot 3</th>
			<th>Slot 4</th>
			<th>Slot 5</th>
		    </tr>
		    <tr>
			<td>
			    <div class="acidjs-css3-treeview" id="slot-1">
				<ul>
				    <li>
	      				<input type="checkbox" checked="checked" id="slot-1-select-all-arrow"/><label><input type="checkbox" id="slot-1-select-all" onChange="checkSubBoxes('slot-1-select-all','slot-1-category')"/><span></span></label><label for="slot-1-select-all-arrow"><b>Select All</b></label> 
					{{range $key,$value := .}} 
					<ul>
	      				    <input type="checkbox" id="slot-1-{{.TrimWhiteSpace $key}}-arrow" checked="checked" /><label><input type="checkbox" id="slot-1-{{.TrimWhiteSpace $key}}" class="slot-1-category" onChange="checkSubBoxes('slot-1-{{.TrimWhiteSpace $key}}','slot-1-{{.TrimWhiteSpace $key}}-mb')"/><span></span></label><label for="slot-1-{{.TrimWhiteSpace $key}}-arrow"><b>{{$key}}</b></label>
					    <ul>
						{{range $mb := $value}}
						<li>
	      					    <input type="checkbox" id="slot-1-{{$mb.Id}}-arrow"  checked="checked" /><label><input type="checkbox" name="slot1" value="{{$mb.Id}}" id="slot-1-{{$mb.Id}}" class="slot-1-{{$value.TrimWhiteSpace $key}}-mb" /><span></span></label><label for="slot-1-{{$mb.Id}}-arrow"><img src="{{if $mb.ImgURL}} {{$mb.ImgURL}} {{else}} https://yhs.apsva.us/wp-content/uploads/legacy_assets/yhs/032bde3c5d-status_gray.png {{end}}" /> {{$mb.Description}}</label>
						</li>
						{{end}}
					    </ul>
					</ul>
					{{end}}
				</ul>
			    </div>
			</td>
			<td>
			    <div class="acidjs-css3-treeview" id="slot-2">
				<ul>
				    <li>
	      				<input type="checkbox" checked="checked" id="slot-2-select-all-arrow"/><label><input type="checkbox" id="slot-2-select-all" onChange="checkSubBoxes('slot-2-select-all','slot-2-category')"/><span></span></label><label for="slot-2-select-all-arrow"><b>Select All</b></label> 
					{{range $key,$value := .}} 
					<ul>
	      				    <input type="checkbox" id="slot-2-{{.TrimWhiteSpace $key}}-arrow" checked="checked" /><label><input type="checkbox" id="slot-2-{{.TrimWhiteSpace $key}}" class="slot-2-category" onChange="checkSubBoxes('slot-2-{{.TrimWhiteSpace $key}}','slot-2-{{.TrimWhiteSpace $key}}-mb')"/><span></span></label><label for="slot-2-{{.TrimWhiteSpace $key}}-arrow"><b>{{$key}}</b></label>
					    <ul>
						{{range $mb := $value}}
						<li>
	      					    <input type="checkbox" id="slot-2-{{$mb.Id}}-arrow"  checked="checked" /><label><input type="checkbox" name="slot2" value="{{$mb.Id}}" id="slot-2-{{$mb.Id}}" class="slot-2-{{$value.TrimWhiteSpace $key}}-mb" /><span></span></label><label for="slot-2-{{$mb.Id}}-arrow"><img src="{{if $mb.ImgURL}} {{$mb.ImgURL}} {{else}} https://yhs.apsva.us/wp-content/uploads/legacy_assets/yhs/032bde3c5d-status_gray.png {{end}}" /> {{$mb.Description}}</label>
						</li>
						{{end}}
					    </ul>
					</ul>
					{{end}}
				</ul>
			    </div>

			</td>
			<td>
			    <div class="acidjs-css3-treeview" id="slot-3">
				<ul>
				    <li>
	      				<input type="checkbox" checked="checked" id="slot-3-select-all-arrow"/><label><input type="checkbox" id="slot-3-select-all" onChange="checkSubBoxes('slot-3-select-all','slot-3-category')"/><span></span></label><label for="slot-3-select-all-arrow"><b>Select All</b></label> 
					{{range $key,$value := .}} 
					<ul>
	      				    <input type="checkbox" id="slot-3-{{.TrimWhiteSpace $key}}-arrow" checked="checked" /><label><input type="checkbox" id="slot-3-{{.TrimWhiteSpace $key}}" class="slot-3-category" onChange="checkSubBoxes('slot-3-{{.TrimWhiteSpace $key}}','slot-3-{{.TrimWhiteSpace $key}}-mb')"/><span></span></label><label for="slot-3-{{.TrimWhiteSpace $key}}-arrow"><b>{{$key}}</b></label>
					    <ul>
						{{range $mb := $value}}
						<li>
	      					    <input type="checkbox" id="slot-3-{{$mb.Id}}-arrow"  checked="checked" /><label><input type="checkbox" name="slot3" value="{{$mb.Id}}" id="slot-3-{{$mb.Id}}" class="slot-3-{{$value.TrimWhiteSpace $key}}-mb" /><span></span></label><label for="slot-3-{{$mb.Id}}-arrow"><img src="{{if $mb.ImgURL}} {{$mb.ImgURL}} {{else}} https://yhs.apsva.us/wp-content/uploads/legacy_assets/yhs/032bde3c5d-status_gray.png {{end}}" /> {{$mb.Description}}</label>
						</li>
						{{end}}
					    </ul>
					</ul>
					{{end}}
				</ul>
			    </div>

			</td>
			<td>
			    <div class="acidjs-css3-treeview" id="slot-4">
				<ul>
				    <li>
	      				<input type="checkbox" checked="checked" id="slot-4-select-all-arrow"/><label><input type="checkbox" id="slot-4-select-all" onChange="checkSubBoxes('slot-4-select-all','slot-4-category')"/><span></span></label><label for="slot-4-select-all-arrow"><b>Select All</b></label> 
					{{range $key,$value := .}} 
					<ul>
	      				    <input type="checkbox" id="slot-4-{{.TrimWhiteSpace $key}}-arrow" checked="checked" /><label><input type="checkbox" id="slot-4-{{.TrimWhiteSpace $key}}" class="slot-4-category" onChange="checkSubBoxes('slot-4-{{.TrimWhiteSpace $key}}','slot-4-{{.TrimWhiteSpace $key}}-mb')"/><span></span></label><label for="slot-4-{{.TrimWhiteSpace $key}}-arrow"><b>{{$key}}</b></label>
					    <ul>
						{{range $mb := $value}}
						<li>
	      					    <input type="checkbox" id="slot-4-{{$mb.Id}}-arrow"  checked="checked" /><label><input type="checkbox" name="slot4" value="{{$mb.Id}}" id="slot-4-{{$mb.Id}}" class="slot-4-{{$value.TrimWhiteSpace $key}}-mb" /><span></span></label><label for="slot-4-{{$mb.Id}}-arrow"><img src="{{if $mb.ImgURL}} {{$mb.ImgURL}} {{else}} https://yhs.apsva.us/wp-content/uploads/legacy_assets/yhs/032bde3c5d-status_gray.png {{end}}" /> {{$mb.Description}}</label>
						</li>
						{{end}}
					    </ul>
					</ul>
					{{end}}
				</ul>
			    </div>

			</td>
			<td>
			    <div class="acidjs-css3-treeview" id="slot-5">
				<ul>
				    <li>
	      				<input type="checkbox" checked="checked" id="slot-5-select-all-arrow"/><label><input type="checkbox" id="slot-5-select-all" onChange="checkSubBoxes('slot-5-select-all','slot-5-category')"/><span></span></label><label for="slot-5-select-all-arrow"><b>Select All</b></label> 
					{{range $key,$value := .}} 
					<ul>
	      				    <input type="checkbox" id="slot-5-{{.TrimWhiteSpace $key}}-arrow" checked="checked" /><label><input type="checkbox" id="slot-5-{{.TrimWhiteSpace $key}}" class="slot-5-category" onChange="checkSubBoxes('slot-5-{{.TrimWhiteSpace $key}}','slot-5-{{.TrimWhiteSpace $key}}-mb')"/><span></span></label><label for="slot-5-{{.TrimWhiteSpace $key}}-arrow"><b>{{$key}}</b></label>
					    <ul>
						{{range $mb := $value}}
						<li>
	      					    <input type="checkbox" id="slot-5-{{$mb.Id}}-arrow"  checked="checked" /><label><input type="checkbox" name="slot5" value="{{$mb.Id}}" id="slot-5-{{$mb.Id}}" class="slot-5-{{$value.TrimWhiteSpace $key}}-mb" /><span></span></label><label for="slot-5-{{$mb.Id}}-arrow"><img src="{{if $mb.ImgURL}} {{$mb.ImgURL}} {{else}} https://yhs.apsva.us/wp-content/uploads/legacy_assets/yhs/032bde3c5d-status_gray.png {{end}}" /> {{$mb.Description}}</label>
						</li>
						{{end}}
					    </ul>
					</ul>
					{{end}}
				</ul>
			    </div>
			</td>

			</td>
		    </tr>
		    <tr style="border: none;">
			<td style="border: none;">
			    <button type="button" onClick="subSlotForm()" id="slot-submit-button">Submit Slot Choices</button>
			    <!-- <input type="submit" value="Submit Slot Choices" /> -->
			</td>
			<td style="border: none;">
			</td>
			<!-- <td>
			     <button type="button">Save as Preset:</button> 
			     <input type="edit" />
			     </td> -->
		    </tr>
		</form>
		<script>
		 function subSlotForm (){
		     $.ajax({
			 url:"/slotSubmit",
			 type:'post',
			 data:$("#slot-submit-form").serialize(),
			 success:function(){
			     alert("Slot choices submitted");
			 }
		     });
		 }

		</script>

	    </table>
	</div>

	<!-- <div>
	     <table>
	     <form>
	     <tr>
	     <td>
	     <h3>Preset Slot Configurations</h3>
	     </td>
	     </tr>
	     <tr>
	     <td>
	     <div class="acidjs-css3-treeview">
	     <ul>
	     <li>
	     <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0"><b>Select All</b></label>                  
	     <ul>
	     <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">Preset 1 </label><button type="button">Edit</button>
	     <ul>
	     <li>
	     <label>badge 1</label>
	     </li>
	     <li>
	     <label>badge 2</label>
	     </li>
	     </ul>
	     <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">Preset 2 </label><button type="button">Edit</button>
	     <ul>
	     <li>
	     <label>badge 3</label>
	     </li>
	     <li>
	     <label>badge 4</label>
	     </li>
	     </ul>

	     </ul>
	     </ul>
	     </div>
	     </td>
	     </tr>
	     <tr>
	     <td style="border: none;">
	     <input type="submit" value="Load Selected Presets" />
	     </td>
	     </tr>
	     </form>
	     </table>
	     </div>
	   -->
	<br />
	<br />
	<br />

	<div>
	    <form action="/setInterval"  method="post" id="interval-form">
		Randomization interval (minutes): 
		<input type="number" name="interval" min=1 value="1"><br />
		<button type="button" onClick="subIntervalForm()">Submit</button>
	    </form>
	    <script>
	     function subIntervalForm (){
		 $.ajax({
		     url:"/setInterval",
		     type:'post',
		     data:$("#interval-form").serialize(),
		     success:function(){
			 alert("Interval set");
		     }
		 });
	     }

	    </script>

	</div>
	
	<br />
	<br />
	<br />

	<div>
	    <form action="/randomize" method="post" id="random-form">
		<button type="button" onClick="subRandomForm()" >Randomize now</button>
	    </form>
	    <script>
	     function subRandomForm (){
		 $.ajax({
		     url:"/randomize",
		     type:'post',
		     data:$("#random-form").serialize(),
		     success:function(){
			 alert("Randomizing microbadges now");
		     }
		 });
	     }

	    </script>

	</div>

    </body>
    <br />
    <br />
    <br />
    <br />
    <br />
    <footer>
	Copyright (C) 2017 Samuel Allen, Allen Technology Solutions
	<br />
	<a href="http://allentechnology.solutions">allentechnology.solutions</a>
    </footer>
</html>
`
