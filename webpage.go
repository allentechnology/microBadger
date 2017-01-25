package main

const htmlPage = `
<html>
  <head>
    <title>MicroBadger</title>
    <style>
      table {
      width: 100%;
      }

      .acidjs-css3-treeview{
      overflow-y:scroll;
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
    <script>
      $(".acidjs-css3-treeview").delegate("label input:checkbox", "change", function() {
    var
        checkbox = $(this),
        nestedList = checkbox.parent().next().next(),
        selectNestedListCheckbox = nestedList.find("label:not([for]) input:checkbox");
 
    if(checkbox.is(":checked")) {
        return selectNestedListCheckbox.prop("checked", true);
    }
    selectNestedListCheckbox.prop("checked", false);
});
    </script>
  </head>
  <body>
    <h1>MicroBadger</h1>
    <div>
      <form>
	Username: 
	<input type="text" />
	<br />
	Password: 
	<input type="password" />
	<br />
	<input type="submit" value="Save Login" />
      </form>
    </div>
    <div id="slots">
      <table>
	<form>
	  <tr>
	    <th>Slot 1</th>
	    <th>Slot 2</th>
	    <th>Slot 3</th>
	    <th>Slot 4</th>
	    <th>Slot 5</th>
	  </tr>
	  <tr>
	    <td>




	      <div class="acidjs-css3-treeview">
		<ul>
		  <li>
		    <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">Select All</label>                  
                  <ul>
                    <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">Category 1</label>
                    <ul>
<li>
                    <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">test 1</label>
</li>
<li>
                    <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">test 2</label>
</li>
                    </ul>
                    <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">Category 2</label>
                    <ul>
<li>
                    <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">test 3</label>
</li>
<li>
                    <input type="checkbox" id="node-0" checked="checked" /><label><input type="checkbox" /><span></span></label><label for="node-0">test 4</label>
</li>
                    </ul>

                  </ul>
                </ul>
	      </div>


	      
	      <div id="slot1" class="slotSelection">
		<input type="checkbox">Select All</input><br />
		<input type="checkbox">Category 1</input><br />
		<input type="checkbox">test1</input><br />
		<input type="checkbox">test2</input><br />
		<input type="checkbox">Category 2</input><br />
		<input type="checkbox">test3</input><br />
		<input type="checkbox">Category 3</input><br />
		<input type="checkbox">test4</input><br />
		<input type="checkbox">test5</input><br />
		<input type="checkbox">test6</input><br />
		<input type="checkbox">test7</input><br />
	      </div>
	    </td>
	    <td>
	      <div id="slot2" class="slotSelection">
		<input type="checkbox">Select All</input><br />
		<input type="checkbox">Category 1</input><br />
		<input type="checkbox">test1</input><br />
		<input type="checkbox">test2</input><br />
		<input type="checkbox">Category 2</input><br />
		<input type="checkbox">test3</input><br />
		<input type="checkbox">Category 3</input><br />
		<input type="checkbox">test4</input><br />
		<input type="checkbox">test5</input><br />
		<input type="checkbox">test6</input><br />
		<input type="checkbox">test7</input><br />
	      </div>
	    </td>
	    <td>
	      <div id="slot3" class="slotSelection">
		<input type="checkbox">Select All</input><br />
		<input type="checkbox">Category 1</input><br />
		<input type="checkbox">test1</input><br />
		<input type="checkbox">test2</input><br />
		<input type="checkbox">Category 2</input><br />
		<input type="checkbox">test3</input><br />
		<input type="checkbox">Category 3</input><br />
		<input type="checkbox">test4</input><br />
		<input type="checkbox">test5</input><br />
		<input type="checkbox">test6</input><br />
		<input type="checkbox">test7</input><br />
	      </div>
	    </td>
	    <td>
	      <div id="slot4" class="slotSelection">
		<input type="checkbox">Select All</input><br />
		<input type="checkbox">Category 1</input><br />
		<input type="checkbox">test1</input><br />
		<input type="checkbox">test2</input><br />
		<input type="checkbox">Category 2</input><br />
		<input type="checkbox">test3</input><br />
		<input type="checkbox">Category 3</input><br />
		<input type="checkbox">test4</input><br />
		<input type="checkbox">test5</input><br />
		<input type="checkbox">test6</input><br />
		<input type="checkbox">test7</input><br />
	      </div>
	    </td>
	    <td>
	      <div id="slot5" class="slotSelection">
		<input type="checkbox">Select All</input><br />
		<input type="checkbox">Category 1</input><br />
		<input type="checkbox">test1</input><br />
		<input type="checkbox">test2</input><br />
		<input type="checkbox">Category 2</input><br />
		<input type="checkbox">test3</input><br />
		<input type="checkbox">Category 3</input><br />
		<input type="checkbox">test4</input><br />
		<input type="checkbox">test5</input><br />
		<input type="checkbox">test6</input><br />
		<input type="checkbox">test7</input><br />
	      </div>
	    </td>
	  </tr>
	  <tr style="border: none;">
	    <td style="border: none;">
	      <input type="Submit" value="Submit Slot Choices">
	    </td>
	    <td style="border: none;">
	      <input type="Submit" value="Load Preset">
	    </td>
	    <td style="border: none;">
	      <input type="Submit" value="Save as Preset">
	    </td>
	  </tr>
	</form>
      </table>
    </div>
    <br />
    <br />
    <br />

    <div>
      <form>
	Randomization interval (minutes): 
	<input type="number" min=1><br />
	<input type="submit">
      </form>
    </div>
    <br />
    <br />
    <br />

    <div>
      <form>
	<input type="submit" value="Randomize now">
      </form>
    </div>
    <button type="button" onclick="document.getElementById('demo').innerHTML = Date()">Click Me</button>
    <p id="test1"></p>

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
