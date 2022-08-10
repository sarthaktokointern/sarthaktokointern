Git Internal Architecture :

Introduction : Git is a version control system. Most of the version control systems (vcs) use delta storage system that is storing incremental difference of each file in a commit with every increasing revision.

<img width="786" alt="Screenshot 2022-08-10 at 3 27 40 PM" src="https://user-images.githubusercontent.com/99721005/183873484-c2060f00-0d7c-4c71-9367-daa67cd1640d.png">

Disadvantage of using delta storage system :

Suppose in the above-mentioned diagram if we want to move from C5 revision to C4 revision, we will just unapply the patch of C5 revision. But if we want to move from nth revision to (n-1000)th revision then we will have to unapply 1000 patches which is a very difficult task.

How to overcome the disadvantage of delta storage system :

In order to overcome the disadvantage possessed by the delta storage system, git uses directed acyclic graph (dag).

What is DAG?

DAG is directed acyclic graph which means that each edge is a directed edge and while traversing the graph each vertex can be visited at most one time i.e. there is no cycle present in the graph.

Objects in git :

Objects are basically immutable files saved on disk which hash key as their name.
There are 4 types of objects in git :

1.) Blob : They are basically large binary objects which represent any file added in the git.
They store data in compressed format. Point to be noted is that each time when a file is added with a **new content**, a new blob is created corresponding to that file.

Hands on : 

a.) create a new folder in your home dir : mkdir test

b.) switch to that folder : cd test

c.)  initialise a git repo : git init 

d.) save a new blank file : vim test.txt

e.) add it to git : git add test.txt

f.) switch to .git : cd .git

g.) list all the directories : ls -la

h.) switch to 'objects' directory : cd objects

i.) ls -la

you will see some output like this : 

<img width="742" alt="Screenshot 2022-08-10 at 3 29 44 PM" src="https://user-images.githubusercontent.com/99721005/183873936-4b6ca7f9-c225-403f-86e3-164cae2001e3.png">


Now if we will add some content to the file test.txt and add the changes to git then a new blob will get created corresponding to that.
The file structure will be something like this :

<img width="695" alt="Screenshot 2022-08-10 at 3 30 18 PM" src="https://user-images.githubusercontent.com/99721005/183874050-35afc101-abad-449c-b1ba-43b078cca41a.png">

2.) Tree : if blob is a file, tree is a directory. 

Hands on:

a.) inside the directory test, create another directory : mkdir dir1

b.) cd dir1

c.) vim test.txt

d.) add all the changes using 'git add .' from the test directory 

Note : tree is created when the added changes are committed, but we can see the architecture of the tree which will get formed using the following command: git write-tree

We will get the following architecture : 




Simplified tree diagram :

