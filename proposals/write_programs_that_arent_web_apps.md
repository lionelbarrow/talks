Learning From Programs That Aren't Web Applications

The Rails community has gotten pretty good at thinking about design patterns and best practices for building web applications. But many of us have little to no experience writing other 
types of software. That's too bad -- because many of the best ideas in web development, like the MVC pattern, can be derived from more general ideas in software engineering. Just like 
how it can be hard to understand a theorem after working through a single equation, it can be tough for us to understand these more general ideas when the only programs we ever write are
web apps.

In this talk, we'll look at the architecture of a few different programs -- a compiler, a desktop GUI application, and a source control system. We'll examine the constraints these programs 
face and the goals they have, and try to tease out how this affects they way their code is organized. From there, we'll extract some general principles of software design, which we'll then apply to 
web applications. We'll be specific, concrete, and thoughtful about the lessons we take away from this exercise -- you'll walk away with much more than a vague sense that "modularity is good". 
Finally, we'll return to our familiar web application patterns and use our new wisdom to generate ideas about why they they work well and when they might not apply to our apps.



Lionel Barrow (@lionelbarrow) is a developer at Braintree and a graduate student in computer science at the University of Chicago. At Braintree, he worked on Venmo Touch, 
once of the first highly available web services to be written in Go. A born geek, he fell into programming when he got sick of manually de-duplicating his college's admissions database. 
When not coding, he blogs (in vim) and dabbles in machine learning -- which, come to think of it, actually involves coding.
