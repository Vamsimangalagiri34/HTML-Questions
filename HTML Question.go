1. What is HTML?

inventor of HTML Tim Berners-Lee.
HTML (HyperText Markup Language) is the language used to create the structure of web pages, defining elements like headings, paragraphs, links, images, and more.
HTML uses tags enclosed in angle brackets (< >) to define elements, e.g., <h1> for headings or <p> for paragraphs.
🌐 for Headings -> h1 to h6 1 has high priority
🌐 p->Paragraphs
🌐 link->Links
🌐 img->Images
🌐 Tables
🌐 Forms
It works alongside CSS (for styling) and JavaScript (for interactivity) to build complete web pages.

An element is a complete structure in HTML, including tags and content, like <p>This is a paragraph.</p>. An attribute adds extra information to an element, such as href in <a href="https://example.com">Link</a>. A tag is the part of an element enclosed in angle brackets, like <h1> or </h1>. A property is a CSS style rule defining an elements appearance or behavior, like color: red;.

HTML tags are used to create the structure and content of web pages. Each tag serves a specific purpose and helps organize the content, making it readable for both browsers and developers.


<meta charset="UTF-8"> unicode transformation format 

Here's how it works in detail:
UTF-8 Encoding:

UTF-8 is a variable-length character encoding. It can represent characters from the entire Unicode character set. This includes everything from ASCII characters (like A, B, C, etc.) to characters from other languages (like é, ñ, 我), and even emojis (like 😊).
UTF-8 uses 1 to 4 bytes to represent a character, which allows it to cover a very wide range of characters.
Browser’s Role:

When the browser reads the document, it looks at the <meta charset="UTF-8"> tag to understand how to decode the file's bytes.
If the file contains non-ASCII characters (such as accented characters or symbols), the browser uses UTF-8 to interpret these characters and display them correctly.
Without UTF-8 (or Incorrect Encoding):

If you don’t specify the correct character encoding, or if there’s an encoding mismatch, the browser might not interpret the bytes correctly. This could result in garbled or broken characters (like question marks or strange symbols) instead of the intended characters.

In Summary:
UTF-8 is used to encode the characters in an HTML page (or any text file) so that the browser can understand and display them correctly.
When a webpage is served to the browser, the text (including special characters) is encoded using UTF-8, ensuring that the browser can decode and display the correct characters on the screen.



2. How to insert an image in HTML?
<img> tag is used to add an image in a web page.
Images are not inserted into a web page basically they are linked to web pages. The <img> tag helps to create a holding space for the referenced image.

The <img> tag is normally empty, it has attributes only, and does not have a closing tag. this type called void tag

 <img> tag has two required parameters:

src – The path to the image
alt – An alternate text for the image

<img src="image path" alt="Italian Trulli">
<img src="demo.jpg" alt="Italian Trulli">

3. How to set background image in HTML?
In order to add a background image on an HTML element you need to use  two things:

the HTML style attribute and 
the CSS background-image property:
<div style="background-image: url('img_boy.jpg');">

4. How to comment in HTML?
Normally HTML comments are not being displayed in the browser. But these comments can help to document the HTML source code.

<!– Write your comments here –>

5. `How to give space in HTML?`
In order to add a space in the webpage, Go  where you want to add the space and then use the spacebar. Normally, HTML displays one space between words, no matter how many times you have entered  the space bar. 

Now if you Type &nbsp; to force an extra space. 

This is known as a non-breaking space because it helps to prevent a line break at its location.

6.`How to link CSS to HTML?`

Types of CSS:

So there are three ways to add CSS to HTML documents :

Inline – by putting the style attribute inside HTML elements

<h1 style="color:red;">A Blue Heading</h1>
 
<p style="color:blue;">A red paragraph.</p>

Internal – by putting  a <style> element in the <head> section

<!DOCTYPE html>
<html>
<head>
<style>
body {background-color: powderblue;}
h1   {color: blue;}
p    {color: red;}
</style>
</head>
<body>

External – by adding a <link> element to link to an external CSS file

<link rel="stylesheet" href="styles.css">

7. How to align text in HTML?

div.a {
	text-align: center;
  }
   
  div.b {
	text-align: left;
  }
   
  div.c {
	text-align: right;
  }
   
  div.c {
	text-align: justify;
  }

8. How to create a table in HTML?

HTML tables help web developers to set the data into rows and columns.

The <table> tag is there in the HTML table.
Each table row can be defined with a <tr> tag.
Each header can be defined with a <th> tag. 
Each data or the cell is defined with a <td> tag.
If your text is in the  <th> elements then they will be bold and centered.
If your text is in the <td> elements then they will be regular and left-aligned.

refer table.html file

11. How to change font color in HTML?
<font> tag, is used to specify the text color.

<font Color=”Blue”>  
<font color=”rgb(128,128,0)”
<font color=”#00FF00″>


13. What is doctype in HTML?
It is a way to give  “information” to the browser about what will be the document type to expect.

A document type declaration, or DOCTYPE, is a line of code that tells a web browser or email client which version of HTML or SGML a document is written in
The <!DOCTYPE> declaration in HTML is used to specify the document type and version of HTML that the page is using.

15. How to add space using < pre > tag in HTML?

The <pre> tag in HTML preserves both spaces and line breaks exactly as written in the HTML code. To add spaces using the <pre> tag, simply include the desired number of spaces directly in the content inside the tag.

Use &nbsp; when you want to add a specific number of spaces between words.
Use the <pre> tag to preserve the exact spacing and formatting of your text block, including line breaks and multiple spaces.

16. What is dom in HTML?


DOM stands for Document Object Model. When a web page is getting loaded that time the browser creates a Document Object Model of the page and it is constructed as a tree of Objects. HTML DOM is basically an Object Model for HTML

The DOM (Document Object Model) in HTML is a programming interface that represents the structure of an HTML document as a tree-like hierarchy. It allows developers to access, manipulate, and modify the content, structure, and styles of a webpage dynamically using JavaScript or other programming languages.

Key Points:
The DOM treats each element, attribute, and piece of content in the HTML document as an object or node.
It creates a tree-like structure where:
The root node is the <html> element.
Child nodes include <head>, <body>, and other nested elements.



18. How to create form in HTML?

check Form.html

Features of the Form:
Text Input: Name, email, and password fields with placeholders.
Radio Buttons: To select gender (only one option can be chosen).
Checkboxes: For hobbies (multiple options can be selected).
Dropdown Menu: To select the country.
Buttons:
Submit: Sends the form data to the specified action URL.
Reset: Clears all inputs in the form.

In HTML, a label is an element used to define a user-readable description for form controls, such as input fields, checkboxes, and radio buttons. It improves accessibility and makes the form easier to use, especially for screen readers.

Key Features:
The <label> element is associated with a specific form control via the for attribute.
The for attribute matches the id of the form control, linking the label to that element.
When you click on the label, it focuses the associated form control (e.g., clicking a label for a text box will focus the input field).

19. How to create button in HTML?
<button type="button">Click Here!</button>
 < input type='button' />

23. How to use div tag in html to divide the page?

In HTML, the <div> tag is a block-level element used to group and organize content into sections. It does not have any inherent styling, but it is commonly used with CSS to divide and structure a webpage into different sections or layouts.
Block-Level Element: A <div> element occupies the full width of its parent container and creates a new line before and after itself.
Grouping Similar Elements: You can wrap multiple elements inside a <div> to group them together. This helps in structuring the page, and you can apply common styles to the grouped elements.
CSS Application: By using the class or id attributes, you can target the <div> element and apply CSS rules to style the grouped content.

30. Which is better: HTML or HTML5?
HTML5 is the newest version of HTML and it is better than HTML because it includes new features like audio and video elements, new semantic elements, and support for local storage.

31 .Semantic elements in HTML are tags that provide meaning about the content they enclose, making the structure of the web page more descriptive and easier to understand for both developers and machines (like search engines and screen readers).
Improved Accessibility: Helps screen readers and assistive technologies understand the content structure.
SEO (Search Engine Optimization): Search engines use semantic tags to understand the importance and context of the content.
Better Readability and Maintainability: Semantic elements improve the clarity of code, making it easier for developers to understand the page structure.

Explanation:
The <header> contains the main heading and navigation links.
The <main> element encloses the main content of the page, which includes a <section> and an <article>.
The <footer> contains the copyright information.

32. What is span in HTML?
The <span> tag in HTML is an inline element used to group and style a small section of text or other inline elements within a block of content. It doesn't inherently add any styling or structure by itself but serves as a container for applying styles or JavaScript functionality to a specific part of the content.

Key Points about <span>:
Inline Element: Unlike block-level elements like <div>, <span> does not break the flow of content. It stays within the same line as the content around it.
Styling and Grouping: Its often used to apply styles (like color, font-size) or JavaScript functionality to a specific portion of text or content within a larger block.
No Semantics: The <span> tag does not convey any particular meaning about the content it wraps. It's purely used for styling and grouping.

The <span> tag is used to apply the class highlight to part of the text ("highlighted text").
The class is defined in CSS to change the text color to red and make it bold.
The rest of the content (the paragraph) remains unaffected because <span> is an inline element. 


When to Use <span>:
When you need to apply styles or JavaScript actions to a small part of a block of text or content, without affecting the layout of the surrounding content.


33. How to underline text in HTML?
<u> tag is used for underline the text. The <u> tag was deprecated in HTML, but then they re-introduced in HTML5.


34. What is a “fieldset” tag in HTML?
A fieldset is used to group related elements in a form. It is useful for creating structures such as tables or grid layouts.

Explanation:
The <fieldset> tag groups the form fields into two sections: Personal Information and Preferences.
The <legend> tag provides a title for each section to describe its content, making the form more understandable.
The browser automatically adds a border around the grouped fields within the <fieldset>, but you can also customize this with CSS.
When to Use <fieldset>:
Use <fieldset> to group related form elements together, especially when creating complex forms. It helps organize your form content and makes it easier for users to navigate and understand the form's structure.

37. How to add a link in HTML? .
To add links in html we use <a> and </a> tags,  which are the tags used to define the links. The <a> tag indicates where the hyperlink starts and the </a> tag indicates where it ends. Whatever text gets added inside these tags, will work as a hyperlink. Add the URL for the link in the <a href=” ”>.

The href (Hypertext Reference) attribute in HTML is used to specify the URL (Uniform Resource Locator) or destination address for an anchor element (<a>). It defines where the link should navigate when clicked.

'''
39. How to create a checkbox in HTML?

check checkbox.html

41. How to add a scroll bar in HTML?

check scrollbar.html

42. What is an attribute in HTML?
HTML attributes help to provide the additional information about HTML elements.

All HTML elements always have attributes
Attributes provide additional information about elements
Attributes always have the start tag
Attributes usually use in name/value pairs like: name=”value”

46. How to bold text in HTML?
To text bold in HTML, use the <b>  </b> tag or <strong> </strong> tag.

<b> hey, welcome to great learning!</b>

Semantic Meaning:
<b> is purely for visual styling (no implied importance).
<strong> conveys that the content is of greater importance or has strong emphasis.
Accessibility: <strong> is more beneficial for accessibility because screen readers will typically emphasize the text when reading it aloud, whereas <b> has no such significance.

The first sentence uses <b>, which only makes the text bold visually.
The second sentence uses <strong>, which makes the text bold and indicates that it has strong emphasis or importance.

55. How to add audio and video ?

<video width="320" height="240" controls>
  <source src="movie.mp4" type="video/mp4">
  Your browser does not support the video tag.
</video>

Attributes:
width and height: Specify the dimensions of the video player.
controls: Adds video controls like play, pause, volume, etc.
src: Specifies the path to the video file.
type: Specifies the MIME type of the video file (e.g., video/mp4, video/webm).


<audio controls>
  <source src="audio.mp3" type="audio/mp3">
  Your browser does not support the audio element.
</audio>


<audio controls autoplay loop muted>
///
Attributes:
controls: Adds audio controls like play, pause, volume, etc.
src: Specifies the path to the audio file.
type: Specifies the MIME type of the audio file (e.g., audio/mp3, audio/ogg).
autoplay: Automatically starts playing the audio as soon as its loaded.
loop: Makes the audio repeat continuously.
muted: Mutes the audio by default.

56. How to add favicon in HTML?

<head>
  <link rel="icon" href="favicon.ico" type="image/x-icon">
</head>

<head>
  <link rel="icon" href="favicon.ico" type="image/x-icon">
</head>

This method will ensure that the favicon is displayed in the browser tab when users visit your site. You can also include different sizes of the favicon for different devices by specifying the sizes attribute.

57. How to embed YouTube video in HTML?
On a computer, go to the YouTube video you want to embed.
Under the video, click SHARE .
Click Embed.
From the box that appears, copy the HTML code.
Paste the code into your blog or website HTML.


Purpose of Metadata:
Improved Search Engine Optimization (SEO): Metadata such as descriptions and keywords can influence how search engines rank and display your pages.
Better User Experience: Metadata like viewport settings ensures your site looks good on all devices.
Social Media Sharing: Open Graph tags help control how your content is displayed when shared on social media platforms.
<meta name="author" content="John Doe">
<meta name="description" content="This is a description of the page.">


what is <i> and <em> ?

y Differences:
Semantic Meaning:

<i> is purely for styling and has no inherent meaning beyond the italics.
<em> conveys emphasis and is semantically meaningful, indicating the text should be stressed.
Accessibility:

<i> does not provide any specific meaning for screen readers or search engines.
<em> is recognized by screen readers and search engines as important or emphasized content.


75. Is it possible to change the color of the bullet?

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Custom Bullet Color</title>
    <style>
        ul {
            list-style-type: none; /* Remove default bullets */
        }

        li::before {
            content: "•"; /* Custom bullet */
            color: green; /* Bullet color */
            font-size: 20px; /* Bullet size */
            margin-right: 10px; /* Spacing between bullet and text */
        }
    </style>
</head>
<body>

    <ul>
        <li>Item 1</li>
        <li>Item 2</li>
        <li>Item 3</li>
    </ul>

</body>
</html>


77. What are Forms in HTML?
If you want to collect the information of the visitors to the webpage, you can add a form to the webpage. Once the user enters the information into the form fields, it is added to a database specified by you.


79. What is a marquee?
A scrolling text that can go in a specific direction across the screen i.e. left, right, up, or down, automatically. For this you can use the tag <marquee> Text to scroll </marquee>.


.scrolling-text {
           
  animation: scroll-text 10s linear infinite;
}

or 


<marquee 
    direction="right" 
    behavior="alternate" 
    scrollamount="10" 
    bgcolor="yellow" 
    height="50px" 
    width="100%" 
    loop="3">
  Hello, this is a scrolling text with multiple attributes!
</marquee>

What is an image map?

image map groups these clickable regions and links within a single image, making it easier to navigate through an image-based interface

An image map groups multiple links within specific, clickable areas of an image. Instead of having multiple individual images or buttons, an image map allows you to define different regions within a single image, each of which can link to different destinations.

followed by 
<map name=""> 
<area shape='' href='' coords="34,44,270,350">
 <map/>

 What is datalist tag?

 <!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Datalist Example</title>
</head>
<body>

<h1>Choose a fruit</h1>

<input list="fruits" name="fruit" id="fruit">
<datalist id="fruits">
  <option value="Apple">
  <option value="Banana"> //id and list values are must same
  <option value="Orange">
  <option value="Mango">
  <option value="Grapes">
</datalist>

</body>
</html>


The <datalist> tag is especially useful when you want to provide suggestions to users but still allow them to input their own values. For example, in forms where users need to select a city, you can provide a list of popular cities but still allow them to type in a city that isnt on the list.

83. What is difference between HTML and XHTML?

. Strictness:
HTML is more flexible. You can skip some things like closing tags, and it still works.
XHTML is stricter. Every tag must be closed, and everything must be written in lowercase.
2. Error Handling:
In HTML, if you make a mistake, the browser will try to fix it and still show the page (but it might not look perfect).
In XHTML, if there's an error, the page won't show at all.
3. Tag Closure:
HTML doesnt always require you to close tags (like <img> or <br>).
XHTML requires closing tags for everything (e.g., <img /> and <br />).

What is the ‘class’ attribute and id in HTML? 

class Attribute:
Purpose: The class attribute is used to group multiple elements together that share the same style or behavior.

Multiple Elements: You can assign the same class to many elements on the page.

Styling: It is typically used for applying CSS styles or targeting elements with JavaScript.


Using the same id for multiple elements in HTML can work, but it's not recommended because id is meant to be unique for each element. Here's why:

1. Uniqueness Requirement:
The id attribute is intended to identify a single unique element on the page. According to the HTML specification, id should be unique within a document.

JavaScript and CSS Behavior:
CSS: If you use the same id for multiple elements, CSS will still apply the styles to all those elements, but it will only properly target the last element with that id. This happens because CSS selectors expect ids to be unique, and the browser may prioritize the last id occurrence.

JavaScript: When you use JavaScript (e.g., document.getElementById()), it returns only the first element with the given id. This means that if you have multiple elements with the same id, JavaScript will only interact with the first one, and you could unintentionally miss the others.

85. What is the use of an IFrame tag?
IFrame or Inline Frame is basically an HTML document implanted inside the other  HTML documents on a website. The IFrame element is used for inserting content from other source, which can be an advertisement into a webpage


86. What is the use of figure tag in HTML 5?
The <figure> tag is used to wrap content like images, videos, or other media along with their caption, making the content more semantic and accessible. It helps to group media elements and gives them a clear context in the document.


87. Why is a URL encoded in HTML?
URL is encoded in HTML as it converts characters into a format that can be transmitted over the web. A URL is transmitted over the internet through the ASCII character set. The non-ASCII characters can be replaced by “%” which is followed by hexadecimal digits.

Examples of ASCII Characters:
Letters: A, B, C, a, b, c
Numbers: 0, 1, 2, 3
Punctuation: !, ., ?, ,
Control Characters: \n, \t

Examples of Non-ASCII Characters:
Accented letters: é, ç, ñ, ö
Other languages: ありがとう (Japanese), 你好 (Chinese)
Special symbols: ©, €, ™
Emojis: 😊, ❤️



90. What is the advantage of collapsing white space?
White spaces are a sequence of blank space characters, treated as a single space character in an HTML document.

The advantages of collapsing white space are 
Helps the content of the code to be more understandable and readable to the users.  
Decreases the transmission time between the server and the client and removes unnecessary bytes that are occupied by the white spaces.
If you leave extra white space by mistake, the browser will ignore it and display the content properly.
90. What is the advantage of collapsing white space?
White spaces are a sequence of blank space characters, treated as a single space character in an HTML document.

The advantages of collapsing white space are 
Helps the content of the code to be more understandable and readable to the users.  
Decreases the transmission time between the server and the client and removes unnecessary bytes that are occupied by the white spaces.
If you leave extra white space by mistake, the browser will ignore it and display the content properly..

92. Describe the HTML layout structure.

<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>HTML Layout</title>
</head>
<body>
  <header>
    <h1>My Website</h1>
    <nav>
      <ul>
        <li><a href="#home">Home</a></li>
        <li><a href="#about">About</a></li>
        <li><a href="#contact">Contact</a></li>
      </ul>
    </nav>
  </header>

  <section>
    <h2>Welcome</h2>
    <p>This is the main content of the page.</p>
  </section>

  <article>
    <h2>Blog Post</h2>
    <p>This is a sample blog post.</p>
  </article>

  <aside>
    <h3>Related Links</h3>
    <p>Check out these links.</p>
  </aside>

  <footer>
    <p>© 2024 My Website</p>
  </footer>
</body>
</html>


<aside> Tag
The <aside> tag is used to define content that is indirectly related to the main content of the page. It is often used for sidebars, advertisements, or additional information that supports the primary content but isn’t essential to it.

Key Points:
Typically used for sidebars, pull quotes, or related links.
Content in <aside> is complementary, not the focus.
Can be placed inside or outside <section> or <article> tags.

<article> Tag
The <article> tag is used to define self-contained content that can be independently distributed or reused. It is commonly used for blog posts, news articles, or user comments.

Key Points:
Content in <article> should make sense on its own.
Often used in blogs, forums, or news sites.
Can contain headings, paragraphs, images, and other HTML elements.

