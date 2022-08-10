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

<img width="693" alt="Screenshot 2022-08-10 at 3 46 12 PM" src="https://user-images.githubusercontent.com/99721005/183877293-fd5133da-eaf0-4952-ae42-da0e8bd18c18.png">



Simplified tree diagram :

<img width="532" alt="Screenshot 2022-08-10 at 3 46 33 PM" src="https://user-images.githubusercontent.com/99721005/183877361-7c4ef4af-bc3c-4320-b5d8-33b4e2a0be3e.png">


3.) Commit : it is an object in git, and it has 2 pointers

a.) First pointer points to the root of the tree

b.) Second pointer points to the previous commit object

Note : the initial commit has only a single pointer pointing towards the root of the tree. It is also known as the head.

Example :

<img width="830" alt="Screenshot 2022-08-10 at 4 08 49 PM" src="https://user-images.githubusercontent.com/99721005/183881771-e823a233-ccb2-449f-a27d-ea535e8e58a5.png">


In above diagram you can easily see :

a.) Red block representing how commits form linked list. 

b.) Green borders showing new root tree formed. New tree is formed if blob corresponding to it has changed. 

c.) Blue arrow shows how directed acyclic graph is formed here.

d.) Dashed orange lines show how git render the repository snapshot when pointed to a commit just by performing Depth First Search (DFS).

4.) Tags : these are just pointers to a commit. So when we checkout to a tag, the git simply performs dfs from the commit to which the tag is pointing and renders the project structure.

Directed Acyclic Graph (DAG) in GIT :

<img width="912" alt="Screenshot 2022-08-10 at 4 30 20 PM" src="https://user-images.githubusercontent.com/99721005/183885450-2b1c4a45-4598-4607-a8df-7af917d5ea94.png">



1. In commit C1; 2 files b1 and b2 are committed with b2 in t2 directory.
2. In commit C2; b1 is changed to b1' and new file b3 is added.
3. In commit C3; a new file b4 is added in t2 directory which changed tree corresponding to t2 directory to t2'.
4. In commit C4 file b1' is changed back to earlier implementation i.e. b1.

Points to notice:
1. With each commit we are getting new root tree.
2. The blobs are created based on content, i.e. if blob corresponding to content of file exist; so tree will point to same blob; which happened in c4; t5 points to b1.
3. You can see this is a complex graph but no where it forms cycle if we move in direction of edges.
4. DAG cycle consist of commit, tree and blob or tree and blob only. Examples c1-t1-t2-t3-c2, t3-t2-b2-t2'-t4-b3, etc.
