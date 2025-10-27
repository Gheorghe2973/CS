# Hello Hacker

## Creating Directories and Copying Files

### Task

Create a directory named "leap" in the home directory and copy `/challenge/secret` into it.

### Solution Steps

1. Create the directory "leap":

```bash
   mkdir ~/leap
```

2. Copy the secret file to the new directory:

```bash
   cp /challenge/secret ~/leap/
```

3. Verify the file was copied correctly:

```bash
   ls -la ~/leap/secret
```

4. Run the challenge to obtain the flag:

```bash
   /challenge/solve
```

### Result

The file `/home/hacker/leap/secret` now exists with the same contents as `/challenge/secret`, and the challenge script provides the flag.

---

## Invoking Your First Command

### Task

Invoke the `hello` command to retrieve the flag and understand basic command execution in Linux.

### Concept

When you type a command and press Enter in a Linux terminal, the shell executes that command. Commands in Linux are case-sensitive, meaning `hello` is different from `HELLO`.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Execute the `hello` command:

```bash
   hello
```

### Result

The `hello` command executes successfully and displays the flag. This demonstrates the basic principle of command invocation in a Linux shell.

---

## Intro to Arguments

### Task

Execute the `hello` command with a single argument `hackers` to retrieve the flag.

### Concept

Arguments are additional data passed to a command. When you type a command line and press Enter, the shell parses your input into:

- The **command** (first word)
- The **arguments** (subsequent words)

For example, in `echo Hello Hackers!`:

- Command: `echo`
- Arguments: `Hello` and `Hackers!`

The `echo` command simply echoes (prints) all of its arguments back to the terminal.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Execute the `hello` command with the argument `hackers`:

```bash
   hello hackers
```

### Result

The `hello` command accepts the argument and displays the flag, demonstrating how commands process arguments passed to them.

---

## Command History

### Task

Use the shell's command history feature to retrieve a flag that has been injected into the history.

### Concept

The shell saves a history of every command you invoke. You can navigate through this saved history using:

- **Up/Down arrow keys**: Scroll through previous commands
- **`history` command**: Display all saved commands

This feature is useful for:

- Repeating previous commands without retyping
- Finding commands you've run before
- Reviewing your command-line activity

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Access command history using one of these methods:

   **Method 1**: Press the up arrow key repeatedly to scroll through history

   **Method 2**: View all history at once:

```bash
   history
```

3. Locate the flag in the command history output.

### Result

The flag is found in the command history, demonstrating how the shell maintains a record of all executed commands that can be accessed and reused.




# Pondering Paths

## The Root

### Task

Invoke the `pwn` program located in the root directory (`/`) using its absolute path to retrieve the flag.

### Concept

The filesystem starts at `/`, called the root directory. Under it are directories, configuration files, programs, and flags.

An **absolute path** is a path that starts with the root directory `/`. For example, `/pwn` is an absolute path that points to the `pwn` program in the root directory.

You can invoke a program by providing its path on the command line. When you specify the exact path starting from `/`, you're using an absolute path.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Invoke the `pwn` program using its absolute path:

```bash
   /pwn
```

### Result

The `pwn` program executes and displays the flag, demonstrating how to invoke programs using absolute paths in the Linux filesystem.


## Program and Absolute Paths

### Task

Execute the `run` program located in the `/challenge` directory using its absolute path to retrieve the flag.

### Concept

Challenges in pwn.college are located in the `challenge` directory, which is in the root directory (`/`). The path to the challenge directory is `/challenge`.

In this level, the challenge program is named `run` and it lives in the `/challenge` directory. Thus, the full absolute path to the `run` challenge program is `/challenge/run`.

To execute a program, you invoke it by providing its absolute path on the command line.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Invoke the `run` program using its absolute path:

```bash
   /challenge/run
```

### Result

The `run` program executes from the `/challenge` directory and displays the flag, demonstrating how to invoke programs using their full absolute paths in nested directories.



## Position Thy Self

### Task

Execute the `/challenge/run` program from a specific directory by first navigating to that directory using the `cd` command.

### Concept

The Linux filesystem has tons of directories with tons of files. You can navigate around directories by using the `cd` (change directory) command and passing a path to it as an argument.

For example:

```bash
cd /some/new/directory
```

This affects the "current working directory" of your process (in this case, the bash shell). Each process has a directory in which it's currently hanging out, and the `cd` command changes this location.

The `~` symbol in the prompt shows the current path that your shell is located at.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Run the challenge program to see which directory it requires:

```bash
   /challenge/run
```

3. The program will tell you which directory to navigate to. Use `cd` to change to that directory:

```bash
   cd /path/to/required/directory
```

(Replace `/path/to/required/directory` with the actual path the program specifies)

4. Execute the challenge program again from the new directory:

```bash
   /challenge/run
```

### Result

By navigating to the correct directory using `cd` before running the challenge program, the flag is successfully retrieved. This demonstrates how the current working directory affects program execution


## Position Elsewhere

### Task

Execute the `/challenge/run` program from a specific directory by navigating to that directory using the `cd` command.

### Concept

You can navigate around directories by using the `cd` (change directory) command and passing a path to it as an argument:

```bash
cd /some/new/directory
```

This affects the "current working directory" of your process. Each process has a directory in which it's currently hanging out, and the `cd` command changes this location.

The `~` symbol in the prompt shows the current path that your shell is located at.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Run the challenge program to see which directory it requires:

```bash
   /challenge/run
```

3. The program will specify a directory. Use `cd` to navigate to that directory:

```bash
   cd /path/to/required/directory
```

(Replace with the actual path the program specifies)

4. Execute the challenge program again from the correct directory:

```bash
   /challenge/run
```

### Result

By navigating to the specified directory before running the challenge program, the flag is successfully retrieved.


## Position Yet Elsewhere

### Task

Execute the `/challenge/run` program from a specific directory by navigating to that directory using the `cd` command.

### Concept

You can navigate around directories by using the `cd` (change directory) command and passing a path to it as an argument:

```bash
cd /some/new/directory
```

This affects the "current working directory" of your process. Each process has a directory in which it's currently hanging out, and the `cd` command changes this location.

The `~` symbol in the prompt shows the current path that your shell is located at.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Run the challenge program to see which directory it requires:

```bash
   /challenge/run
```

3. The program will specify a directory. Use `cd` to navigate to that directory:

```bash
   cd /path/to/required/directory
```

(Replace with the actual path the program specifies)

4. Execute the challenge program again from the correct directory:

```bash
   /challenge/run
```

### Result

By navigating to the specified directory before running the challenge program, the flag is successfully retrieved


## Implicit Relative Paths, from /

### Task

Execute the `/challenge/run` program using a relative path while having a current working directory of `/`.

### Concept

Now you're familiar with absolute paths and changing directories. If you use absolute paths everywhere, it doesn't matter what directory you are in.

However, the current working directory does matter for **relative paths**:

- A **relative path** is any path that does not start at root (i.e., it does not start with `/`)
- A relative path is interpreted **relative** to your current working directory (cwd)
- Your **cwd** is the directory that your prompt is currently located at

This means how you specify a particular file depends on where the terminal prompt is located.

For example, to access `/tmp/a/b/my_file`:

- If cwd is `/`, then relative path is `tmp/a/b/my_file`
- If cwd is `/tmp`, then relative path is `a/b/my_file`
- If cwd is `/tmp/a/b/c`, then relative path is `../my_file` (where `..` refers to parent directory)

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Navigate to the root directory:

```bash
   cd /
```

3. Run the challenge using a relative path (starting with `c`):

```bash
   challenge/run
```

### Result

By executing the program using a relative path from the root directory, the flag is successfully retrieved. This demonstrates how relative paths are interpreted based on the current working directory.


## Explicit Relative Paths, from /

### Task

Execute the `/challenge/run` program using an explicit relative path with `.` while having a current working directory of `/`.

### Concept

Previously, the relative path was "naked": it directly specified the directory to descend into from the current directory. In this level, we explore more explicit relative paths.

In most operating systems, including Linux, every directory has two implicit entries that you can reference in paths:

- `.` refers to the **same directory** (current directory)
- `..` refers to the **parent directory**

These can be used to create explicit relative paths. For example, the following absolute paths are all identical:

- `/challenge`
- `/challenge/.`
- `/challenge/././././././././`
- `/././challenge/././`

The following relative paths are also all identical to each other (when cwd is `/`):

- `challenge`
- `./challenge`
- `././challenge`
- `challenge/.`

Of course, if your current working directory is `/`, the above relative paths are equivalent to the above absolute paths.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Navigate to the root directory:

```bash
   cd /
```

3. Run the challenge using an explicit relative path with `.`:

```bash
   ./challenge/run
```

### Result

By executing the program using an explicit relative path (with `.` notation) from the root directory, the flag is successfully retrieved. This demonstrates how `.` can be used to explicitly reference the current directory in relative paths.


## Implicit Relative Path

### Task

Execute the `run` program from the `/challenge` directory using an explicit relative path with `.` notation.

### Concept

In this level, we practice referring to paths using `.` a bit more. This challenge requires you to run it from the `/challenge` directory, where things get slightly tricky.

Linux explicitly avoids automatically looking in the current directory when you provide a "naked" path. Consider the following:

```bash
hacker@dojo:~$ cd /challenge
hacker@dojo:/challenge$ run
```

This will **not** invoke `/challenge/run`. This is actually a safety measure: if Linux searched the current directory for programs every time you entered a naked path, you could accidentally execute programs in your current directory that happened to have the same names as core system utilities!

As a result, the above commands will yield the following error:

```bash
bash: run: command not found
```

To explicitly tell Linux that you want to execute a program in the current directory, you use `.` like in the previous levels. The way to do this is to explicitly use relative paths to launch `run` in this scenario, using `.` to reference the current directory.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Navigate to the `/challenge` directory:

```bash
   cd /challenge
```

3. Run the program using an explicit relative path:

```bash
   ./run
```

### Result

By using `./run` to explicitly reference the program in the current directory, the flag is successfully retrieved. This demonstrates the security feature where Linux requires explicit notation to execute programs from the current directory.

# Comprehending Commands

## Home Sweet Home

### Task

Execute `/challenge/run` with an absolute path argument pointing to a file in your home directory, where the argument is three characters or less before expansion.

### Concept

Every user has a **home directory**, typically under `/home` in the filesystem. In the dojo, you are the `hacker` user, and your home directory is `/home/hacker`. The home directory is typically where users store most of their personal files, and as you make your way through pwn.college, this is where you'll store most of your solutions.

Typically, your shell session will start with your home directory as your current working directory. The `~` in the prompt is the current working directory, with `~` being shorthand for `/home/hacker`.

Bash provides and uses this shorthand because most of your time will be spent in your home directory. Thus, whenever bash sees `~` provided as the start of an argument in a way consistent with a path, it will expand it to your home directory.

Examples:

```bash
echo LOOK: ~           # Expands to: /home/hacker
cd ~                   # Changes to /home/hacker
cd ~/asdf              # Changes to /home/hacker/asdf
```

**Important**: The expansion of `~` is an **absolute path**, and only the leading `~` is expanded. This means:

- `~/` will be expanded to `/home/hacker/`
- `~/-` will be expanded to `/home/hacker/-` (not `/home/hacker/home/hacker`)

**Fun fact**: `cd` will use your home directory as the default destination if no path is provided.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Run the challenge with a path argument using `~` (three characters or less):

```bash
   /challenge/run ~/f
```

The argument `~/f` satisfies all constraints:

- It's 3 characters before expansion (`~/f`)
- After expansion, it becomes `/home/hacker/f` (an absolute path)
- The path is inside your home directory

### Result

By using the `~` shorthand to create a short argument that expands to an absolute path in the home directory, the flag is successfully retrieved. This demonstrates bash's home directory expansion feature and path constraints

## cat: not the pet, but the command!

### Task
Use the `cat` command to read the flag file that has been copied to your home directory.

### Concept
One of the most critical Linux commands is `cat`. The `cat` command is most often used for reading out files, like so:
```bash
cat /challenge/DESCRIPTION.md
```

`cat` will concatenate (hence the name) multiple files if provided multiple arguments. For example:
```bash
cat myfile
# Output: This is my file!

cat yourfile
# Output: This is your file!

cat myfile yourfile
# Output: This is my file!
#         This is your file!

cat myfile yourfile myfile
# Output: This is my file!
#         This is your file!
#         This is my file!
```

Finally, if you give no arguments at all, `cat` will read from the terminal input and output it (we'll explore that in later challenges).

In this challenge, the flag is copied to the `flag` file in your home directory (where your shell starts).

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Start the challenge (if needed) - this copies the flag to your home directory

3. Read the flag file using `cat`:
```bash
   cat flag
```
   
   Or with the full path:
```bash
   cat ~/flag
```

### Result
The `cat` command displays the contents of the flag file, successfully retrieving the flag. This demonstrates the basic usage of `cat` for reading file contents.

## Catting Absolute Paths

### Task
Use the `cat` command with an absolute path to read the flag file located at `/flag`.

### Concept
In the last level, you used `cat flag` to read the flag out of your home directory. You can, of course, specify `cat`'s arguments as absolute paths.

In this challenge, the flag will not be copied to your home directory, but it will be made readable. You can read it with `cat` at its absolute path: `/flag`.

**Fun Fact**: `/flag` is where the flag *always* lives in pwn.college, but unlike in this challenge, you typically can't access that file directly.

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Read the flag file using `cat` with the absolute path:
```bash
   cat /flag
```

### Result
The `cat` command reads the flag from its absolute path `/flag`, successfully retrieving the flag. This demonstrates using `cat` with absolute paths to access files anywhere in the filesystem.

## Grepping for a Needle in a Haystack

### Task
Use the `grep` command to search for the flag in a large text file containing hundred thousand lines.

### Concept
Sometimes, the files that you might `cat` out are too big. Luckily, we have the `grep` command to search for the contents we need! We'll learn it in this challenge.

There are many ways to `grep`, and we'll learn one way here:
```bash
grep SEARCH_STRING /path/to/file
```

Invoked like this, `grep` will search the file for lines of text containing `SEARCH_STRING` and print them to the console.

In this challenge, a hundred thousand lines of text have been put into the `/challenge/data.txt` file. You need to `grep` it for the flag!

**HINT**: The flag always starts with the text `pwn.college`.

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Use `grep` to search for "pwn.college" in the data file:
```bash
   grep pwn.college /challenge/data.txt
```

### Result
The `grep` command searches through the hundred thousand lines and finds the line containing "pwn.college", which is the flag. This demonstrates how `grep` efficiently searches for specific text patterns in large files.

## Comparing Files

### Task
Use the `diff` command to compare two files and find the real flag among fake flags.

### Concept
When looking for changes between similar files, eyeballing them might not be the most efficient approach! This is where the `diff` command becomes invaluable.

`diff` compares two files line by line and shows you exactly what's different between them. For example:
```bash
hacker@dojo:~$ cat file1
hello
world
hacker@dojo:~$ cat file2
hello
universe
hacker@dojo:~$ diff file1 file2
2c2
< world
---
> universe
```

The output tells us that line 2 changed (`2c2`), with `world` in the first file (`<`) being replaced by `universe` in the second file (`>`).

Sometimes, when new lines are added, you'll see something like:
```bash
hacker@dojo:~$ cat old
pwn
hacker@dojo:~$ cat new
pwn
college
hacker@dojo:~$ diff old new
1a2
> college
```

This tells us that after line 1 in the first file, the second file has an additional line (`1a2` means "after line 1 of file1, add line 2 of file2").

### Challenge Setup
There are two files in `/challenge`:
- `/challenge/decoys_only.txt` contains 100 fake flags
- `/challenge/decoys_and_real.txt` contains all 100 fake flags plus the one real flag

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Use `diff` to find what's different between the two files:
```bash
   diff /challenge/decoys_only.txt /challenge/decoys_and_real.txt
```

3. The output will show the real flag that appears in the second file but not in the first.

### Result
The `diff` command reveals the difference between the two files - the real flag that was added to the second file. This demonstrates how `diff` efficiently identifies changes between similar files.

## Listing Files

### Task
Use the `ls` command to list files in a directory and find the renamed challenge file.

### Concept
So far, we've told you which files to interact with. But directories can have lots of files (and other directories) inside them, and we won't always be here to tell you their names. You'll need to learn to list their contents using the `ls` command!

`ls` will list files in all the directories provided to it as arguments, and in the current directory if no arguments are provided. Observe:
```bash
hacker@dojo:~$ ls /challenge
run
hacker@dojo:~$ ls
Desktop    Downloads  Pictures  Templates
Documents  Music      Public    Videos
hacker@dojo:~$ ls /home/hacker
Desktop    Downloads  Pictures  Templates
Documents  Music      Public    Videos
hacker@dojo:~$
```

In this challenge, `/challenge/run` has been renamed to some random name! You need to list the files in `/challenge` to find it.

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. List the files in the `/challenge` directory:
```bash
   ls /challenge
```

3. Identify the renamed file from the output (it will be the only file or the one with a random name)

4. Execute the file using its new name:
```bash
   /challenge/RENAMED_FILENAME
```
   (Replace `RENAMED_FILENAME` with the actual filename found)

### Result
By using `ls` to list directory contents, the renamed challenge file is identified and executed successfully, retrieving the flag. This demonstrates how `ls` is essential for discovering files in directories.

## Touching Files

### Task
Use the `touch` command to create two specific files, then run the challenge to get the flag.

### Concept
Of course, you can also *create* files! There are several ways to do this, but we'll look at a simple command here. You can create a new, blank file by *touching* it with the `touch` command:
```bash
hacker@dojo:~$ cd /tmp
hacker@dojo:/tmp$ ls
hacker@dojo:/tmp$ touch pwnfile
hacker@dojo:/tmp$ ls
pwnfile
hacker@dojo:/tmp$
```

It's that simple! In this level, please create two files: `/tmp/pwn` and `/tmp/college`, and run `/challenge/run` to get your flag!

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Create the first file `/tmp/pwn`:
```bash
   touch /tmp/pwn
```

3. Create the second file `/tmp/college`:
```bash
   touch /tmp/college
```

4. Run the challenge to get the flag:
```bash
   /challenge/run
```

### Result
By using `touch` to create the two required files, the challenge verifies their existence and provides the flag. This demonstrates how `touch` creates new empty files.

## Removing Files

### Task
Use the `rm` command to delete a file, then run the check program to verify and get the flag.

### Concept
Files are all around you. Like candy wrappers, they'll eventually be too many of them. In this level, we'll learn to clean up!

In Linux, you remove files with the `rm` command, as so:
```bash
hacker@dojo:~$ touch PWN
hacker@dojo:~$ touch COLLEGE
hacker@dojo:~$ ls
COLLEGE    PWN
hacker@dojo:~$ rm PWN
hacker@dojo:~$ ls
COLLEGE
hacker@dojo:~$
```

Let's practice. This challenge will create a `delete_me` file in your home directory! Delete it, then run `/challenge/check`, which will make sure you've deleted it and then give you the flag!

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Start the challenge (if needed) to create the `delete_me` file

3. Delete the file using `rm`:
```bash
   rm ~/delete_me
```
   
   Or simply:
```bash
   rm delete_me
```
   (if you're in the home directory)

4. Run the check program to verify and get the flag:
```bash
   /challenge/check
```

### Result
By using `rm` to delete the specified file, the check program verifies the deletion and provides the flag. This demonstrates how to remove unwanted files from the filesystem.

## Moving Files

### Task
Use the `mv` command to move the `/flag` file to a new location, then run the check program to get the flag.

### Concept
You can also *move* files around with the `mv` command. The usage is simple:
```bash
hacker@dojo:~$ ls
my-file
hacker@dojo:~$ cat my-file
PWN!
hacker@dojo:~$ mv my-file your-file
hacker@dojo:~$ ls
your-file
hacker@dojo:~$ cat your-file
PWN!
hacker@dojo:~$
```

This challenge wants you to move the `/flag` file into `/tmp/hack-the-planet` (do it!). When you're done, run `/challenge/check`, which will check things out and give the flag to you.

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Move the `/flag` file to `/tmp/hack-the-planet`:
```bash
   mv /flag /tmp/hack-the-planet
```

3. Run the check program to verify and get the flag:
```bash
   /challenge/check
```

### Result
By using `mv` to move the flag file to the specified location, the check program verifies the move and provides the flag (by reading from the new location). This demonstrates how to relocate files in the filesystem.

## Hidden Files

### Task
Use the `ls -a` command to find and read a hidden flag file in the root directory.

### Concept
Interestingly, `ls` doesn't list *all* the files by default. Linux has a convention where files that start with a `.` don't show up by default in `ls` and in a few other contexts. To view them with `ls`, you need to invoke `ls` with the `-a` flag, as so:
```bash
hacker@dojo:~$ touch pwn
hacker@dojo:~$ touch .college
hacker@dojo:~$ ls
pwn
hacker@dojo:~$ ls -a
.college    pwn
hacker@dojo:~$
```

Now, it's your turn! Go find the flag, hidden as a dot-prepended file in `/`.

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. List all files (including hidden ones) in the root directory:
```bash
   ls -a /
```

3. Look for a hidden file that appears to be the flag (likely starting with `.flag` or similar)

4. Read the hidden flag file using `cat`:
```bash
   cat /.flag-XXXXX
```
   (Replace `.flag-XXXXX` with the actual hidden filename found)

### Result
By using `ls -a` to show hidden files, the dot-prepended flag file is discovered and read successfully. This demonstrates how Linux hides files starting with `.` by default and how to reveal them.

## An Epic Filesystem Quest

### Task
Follow a series of clues hidden in files throughout the filesystem to eventually find the flag.

### Concept
With your knowledge of `cd`, `ls`, and `cat`, we're ready to play a little game!

We'll start it out in `/`. Normally, when you navigate to `/` and list its contents, you'll see many directories like:
- `bin`, `challenge`, `etc`, `home`, `lib32`, `lib64`, `mnt`, `proc`, `run`, `srv`, `tmp`, `var`
- `boot`, `dev`, `flag`, `lib`, `lib64`, `media`, `opt`, `root`, `sbin`, `sys`, `usr`

You'll recognize the `flag` file and the `challenge` directory.

In this challenge, the flag has been *hidden*! You will use `ls` and `cat` to follow breadcrumbs and find it! Here's how it'll work:

0. Your first clue is in `/`. Head on over there.
1. Look around with `ls`. There'll be a file named HINT or CLUE or something along those lines!
2. `cat` that file to read the clue!
3. Depending on what the clue says, head on over to the next directory (or don't!).
4. Follow the clues to the flag!

### Solution Steps

1. Navigate to root and find the first clue:
```bash
   cd /
   ls
   cat SECRET
```

2. The clue directed to `/usr/share/X11/locale/vi_VN.tcvn` with a **hidden** file:
```bash
   cd /usr/share/X11/locale/vi_VN.tcvn
   ls -a
   cat .WHISPER
```

3. Next clue at `/opt/linux/linux-5.4/include/config/watchdog/handle` was **trapped** (can't cd into it):
```bash
   ls -a /opt/linux/linux-5.4/include/config/watchdog/handle/
   cat /opt/linux/linux-5.4/include/config/watchdog/handle/ALERT-TRAPPED
```

4. Another **trapped** clue at `/usr/share/javascript/mathjax/jax/output/HTML-CSS/fonts/Neo-Euler/Operators/Regular`:
```bash
   ls -a /usr/share/javascript/mathjax/jax/output/HTML-CSS/fonts/Neo-Euler/Operators/Regular
   cat /usr/share/javascript/mathjax/jax/output/HTML-CSS/fonts/Neo-Euler/Operators/Regular/CUE-TRAPPED
```

5. The clue at `/usr/share/doc/libxi6` was **delayed** (requires cd to become readable):
```bash
   cd /usr/share/doc/libxi6
   ls -a
   cat CLUE
```

6. Next **hidden** clue at `/usr/share/javascript/mathjax/unpacked/jax/input/TeX`:
```bash
   ls -a /usr/share/javascript/mathjax/unpacked/jax/input/TeX
   cat /usr/share/javascript/mathjax/unpacked/jax/input/TeX/.DOSSIER
```

7. Another **trapped** clue at `/var/cache/man/fi`:
```bash
   ls -a /var/cache/man/fi
   cat /var/cache/man/fi/LEAD-TRAPPED
```

8. Navigate to `/opt/linux/linux-5.4/drivers/devfreq`:
```bash
   cd /opt/linux/linux-5.4/drivers/devfreq
   ls -a
   cat SPOILER
```

9. Final clue at `/usr/share/javascript/mathjax/localization/fr`:
```bash
   cd /usr/share/javascript/mathjax/localization/fr
   ls -a
   cat SNIPPET
```

### Result
By systematically following the breadcrumb trail through 9 different directories, using various techniques:
- `ls -a` to find hidden files (starting with `.`)
- Reading files without `cd`ing into trapped directories
- Using `cd` when clues were delayed

The flag was successfully located: `pwn.college{A8h20XKva1-oybAfc8a3ji0-fwJ.QX5IDO0wCN4YTOzEzW}`

This demonstrates practical navigation and file reading skills in a complex filesystem exploration scenario with multiple challenges (hidden files, trapped directories, and delayed clues).

## Making Directories

### Task
Use the `mkdir` command to create a directory, then create a file inside it, and run the check program to get the flag.

### Concept
We can create files. How about directories? You make directories using the `mkdir` command. Then you can stick files in there!

Watch:
```bash
hacker@dojo:~$ cd /tmp
hacker@dojo:/tmp$ ls
hacker@dojo:/tmp$ ls
hacker@dojo:/tmp$ mkdir my_directory
hacker@dojo:/tmp$ ls
my_directory
hacker@dojo:/tmp$ cd my_directory
hacker@dojo:/tmp/my_directory$ touch my_file
hacker@dojo:/tmp/my_directory$ ls
my_file
hacker@dojo:/tmp/my_directory$ ls /tmp/my_directory/my_file
/tmp/my_directory/my_file
hacker@dojo:/tmp/my_directory$
```

Now, go forth and create a `/tmp/pwn` directory and make a `college` file in it! Then run `/challenge/run`, which will check your solution and give you the flag!

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Create the `/tmp/pwn` directory:
```bash
   mkdir /tmp/pwn
```

3. Create the `college` file inside the directory:
```bash
   touch /tmp/pwn/college
```

4. Run the challenge to verify and get the flag:
```bash
   /challenge/run
```

### Result
By using `mkdir` to create a directory and `touch` to create a file within it, the challenge verifies the structure and provides the flag. This demonstrates how to create and organize directories and files in the filesystem.

## Finding Files

### Task
Use the `find` command to locate a file named `flag` hidden somewhere in the filesystem, then read it.

### Concept
So now we know how to list, read, and create files. But how do we find them? We use the `find` command!

The `find` command takes optional arguments describing the search criteria and the search location. If you don't specify a search criteria, `find` matches every file. If you don't specify a search location, `find` uses the current working directory (`.`).

For example:
```bash
hacker@dojo:~$ mkdir my_directory
hacker@dojo:~$ mkdir my_directory/my_subdirectory
hacker@dojo:~$ touch my_directory/my_file
hacker@dojo:~$ touch my_directory/my_subdirectory/my_subfile
hacker@dojo:~$ find

./my_directory
./my_directory/my_subdirectory
./my_directory/my_subdirectory/my_subfile
./my_directory/my_file
hacker@dojo:~$
```

And when specifying the search location:
```bash
hacker@dojo:~$ find my_directory/my_subdirectory
my_directory/my_subdirectory
my_directory/my_subdirectory/my_subfile
hacker@dojo:~$
```

We can specify the criteria! For example, filter by name:
```bash
hacker@dojo:~$ find -name my_subfile
./my_directory/my_subdirectory/my_subfile
hacker@dojo:~$ find -name my_subdirectory
./my_directory/my_subdirectory
hacker@dojo:~$
```

You can search the whole filesystem if you want:
```bash
hacker@dojo:~$ find / -name hacker
/home/hacker
hacker@dojo:~$
```

Now it's your turn. The flag has been hidden in a random directory on the filesystem. It's still called `flag`. Go find it!

**Important notes:**
- There are other files named `flag` on the filesystem. Don't panic if the first one you try doesn't have the actual flag in it.
- There are plenty of places in the filesystem that are not accessible to a normal user. These will cause `find` to generate errors, but you can ignore those; we won't hide the flag there!
- `find` can take a while; be patient!

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Use `find` to search the entire filesystem for files named `flag`:
```bash
   find / -name flag
```
   
   (This will search from the root directory `/` for any file with the name `flag`)

3. The command will output multiple locations. Look through the results to find the actual flag file (ignore permission errors)

4. Once you identify the correct location, read the flag:
```bash
   cat /path/to/flag
```
   (Replace with the actual path found)

### Result
By using `find / -name flag` to search the entire filesystem, the hidden flag file is located despite being in a random directory. Reading it with `cat` retrieves the flag. This demonstrates how `find` is a powerful tool for locating files across the entire filesystem by name or other criteria.

## Linking Files

### Task
Create a symbolic link to trick a program into reading the flag from a different location.

### Concept
If you use Linux (or computers) for any reasonable length of time to do any real work, you will eventually run into some variant of the following situation: you want two programs to access the same data, but the programs expect that data to be in two different locations. Luckily, Linux provides a solution to this quandary: **links**.

Links come in two flavors: **hard** and **soft** (also known as **symbolic**) links. We'll differentiate the two with an analogy:

- A **hard link** is when you address your apartment using multiple addresses that all lead directly to the same place (e.g., Apt 2 vs Unit 2)
- A **soft link** is when you move apartments and have the postal service automatically forward your mail from your old place to your new place

In a filesystem, a file is, conceptually, an address at which the contents of that file live. A **hard link** is an alternate address that indexes that data --- accesses to the hard link and accesses to the original file are completely identical, in that they immediately yield the necessary data. A **soft/symbolic link**, instead, contains the original file name. When you access the symbolic link, Linux will realize that it is a symbolic link, read the original file name, and then (typically) automatically access that file. In most cases, both situations result in accessing the original data, but the mechanisms are different.

Hard links sound simpler to most people (case in point, I explained it in one sentence above, versus two for soft links), but they have various downsides and implementation gotchas that make soft/symbolic links, by far, the more popular alternative.

In this challenge, we will learn about symbolic links (also known as **symlinks**). Symbolic links are created with the `ln` command with the `-s` argument, like so:
```bash
hacker@dojo:~$ cat /tmp/myfile
This is my file!
hacker@dojo:~$ ln -s /tmp/myfile /home/hacker/ourfile
hacker@dojo:~$ cat ~/ourfile
This is my file!
hacker@dojo:~$
```

You can see that accessing the symlink results in getting the original file contents! Also, you can see the usage of `ln -s`. Note that the original file path comes **before** the link path in the command!

A symlink can be identified as such with a few methods. For example, the `file` command, which takes a filename and tells you what type of file it is, will recognize symlinks:
```bash
hacker@dojo:~$ file /tmp/myfile
/tmp/myfile: ASCII text
hacker@dojo:~$ file ~/ourfile
/home/hacker/ourfile: symbolic link to /tmp/myfile
hacker@dojo:~$
```

### Challenge Setup
Okay, now you try it! In this level the flag is, as always, in `/flag`, but `/challenge/catflag` will instead read out `/home/hacker/not-the-flag`. Use the symlink, and fool it into giving you the flag!

### Solution Steps

1. Connect to the dojo environment:
```bash
   ssh hacker@dojo.pwn.college
```

2. Create a symbolic link from `/home/hacker/not-the-flag` to `/flag`:
```bash
   ln -s /flag /home/hacker/not-the-flag
```
   
   (Remember: original file path comes first, then the link path)

3. Run the challenge program, which will follow the symlink:
```bash
   /challenge/catflag
```

### Result
By creating a symbolic link that points from `/home/hacker/not-the-flag` to `/flag`, when `/challenge/catflag` tries to read `/home/hacker/not-the-flag`, it follows the symlink and actually reads `/flag`, successfully retrieving the flag. This demonstrates how symbolic links can redirect file access to different locations transparently.