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



## Learning From Documentation

### Task

Read the program's documentation to understand what argument is needed, then run the program with the correct argument to get the flag.

### Concept

The typical need you'll have for documentation is just to figure out how to use all these dang programs, and a specific case of that is figuring out what arguments to specify on the command line. This module will mostly dig into that concept, as a proxy for figuring out how to use the programs in general. Through the rest of the module, you'll go through various ways of asking the environment for help for the programs, but first, we'll dig into the concept of reading documentation.

The correct usage of programs depends, in a large part, on the proper specification of arguments to them. Recall the `-a` of `ls -a` in the **hidden files** challenge of the **Basic Commands** module: that `-a` was an *argument* that told `ls` to list out hidden files as well as non-hidden files. Because we *wanted* to list out hidden files, invoking `ls` with the `-a` argument was the correct way to use it in our scenario.

The program for this challenge is `/challenge/challenge`, and you'll need to invoke it properly in order for it to give you the flag.

### Challenge Documentation

Welcome to the documentation for `/challenge/challenge`! To properly run this program, you will need to pass it the argument of `--giveflag`. Good luck!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Read the documentation (provided in the challenge description) which tells us to use `--giveflag` argument
3. Run the challenge with the correct argument:

```bash
   /challenge/challenge --giveflag
```

### Result

By reading the documentation and understanding that the program requires the `--giveflag` argument, the flag is successfully retrieved. This demonstrates the importance of reading program documentation to understand proper usage and required arguments.



## Learning Complex Usage

### Task

Use a program with a complex argument that itself takes an argument to print the flag file.

### Concept

While using most commands is straightforward, the usage of some commands can get quite complex. For example, the arguments to commands like `sed` and `awk`, which we're definitely not getting into right now, are entire programs in an esoteric programming language! Somewhere on the spectrum between `cd` and `awk` are commands that take arguments to their arguments...

This sounds crazy, but you've already encountered this with the `find` level in **Basic Commands**. `find` has a `-name` argument, and the `-name` argument itself takes an argument specifying the name to search for. Many other commands are analogous.

### Challenge Documentation

Here is this level's documentation for `/challenge/challenge`:

Welcome to the documentation for `/challenge/challenge`! This program allows you to print arbitrary files to the terminal, when given the `--printfile` argument. The argument to the `--printfile` argument is the path of the flag to read. For example, `/challenge/challenge --printfile /challenge/DESCRIPTION.md` will print out the description of the level!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Read the documentation which tells us:

   - Use the `--printfile` argument
   - Pass the path of the file to read as an argument to `--printfile`
   - The flag is at `/flag`
3. Run the challenge with the correct arguments:

```bash
   /challenge/challenge --printfile /flag
```

### Result

By understanding that `--printfile` is an argument that itself takes another argument (the file path), and providing `/flag` as that path, the program prints the flag contents to the terminal. This demonstrates how some commands have complex argument structures where arguments can take their own arguments.



## Reading Manuals

### Task

Use the `man` command to read the manual page for the challenge program and discover the secret option needed to get the flag.

### Concept

This level introduces the `man` command. `man` is short for `manual`, and will display (if available) the manual of the command you pass as an argument. For example, let's say we wanted to learn about the `yes` command (*yes*, this is a real command):

```bash
hacker@dojo:~$ man yes
```

This will display the manual page for `yes`, which will look something like this:

```
YES(1)                          User Commands                         YES(1)

NAME
       yes - output a string repeatedly until killed

SYNOPSIS
       yes [STRING]...
       yes OPTION

DESCRIPTION
       Repeatedly output a line with all specified STRING(s), or 'y'.

       --help display this help and exit

       --version
              output version information and exit
```

### Important Manual Sections

**NAME**: This gives the name (and short description) of the command or concept discussed by the page.

**SYNOPSIS**: This gives a short usage synopsis. These synopses have a standard format. Typically, each line is a different valid invocation of the command, and the lines can be read as:

- `COMMAND [OPTIONAL_ARGUMENT] SINGLE_MANDATORY_ARGUMENT`
- `COMMAND [OPTIONAL_ARGUMENT] MULTIPLE_ARGUMENTS...`

**DESCRIPTION**: Details of the command or concept, with detailed descriptions of the various options.

**SEE ALSO**: Other man pages or online resources that might be useful.

You can scroll around the manpage with your arrow keys and PgUp/PgDn. When you're done reading the manpage, you can hit `q` to quit.

**Note**: Manpages are stored in a centralized database in the `/usr/share/man` directory, but you never need to interact with it directly: you just query it using the `man` command. The arguments to the `man` command aren't file paths, but just the names of the entries themselves (e.g., you run `man yes` to look up the `yes` manpage, rather than `man /usr/bin/yes`, which would be the actual path to the `yes` program but would result in `man` displaying garbage).

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Read the manual page for the challenge program:

```bash
   man challenge
```

3. Scroll through the manual using arrow keys or PgUp/PgDn to find the secret option
4. Look for a description of an argument or option that will print the flag
5. Press `q` to quit the manual when done reading
6. Run the challenge with the discovered argument from the manual:

```bash
   /challenge/challenge [SECRET_OPTION]
```

(Replace `[SECRET_OPTION]` with the actual option found in the manual)

### Result

By using `man challenge` to read the manual page, the secret option is discovered in the documentation. Running the challenge with this option successfully retrieves the flag. This demonstrates the importance of reading manual pages to understand program usage and discover available options.




## Searching Manuals

### Task

Use the search functionality within `man` pages to find the option that will give you the flag.

### Concept

You can scroll man pages with the arrow keys (and PgUp/PgDn) and search with `/`. After searching, you can hit `n` to go to the next result and `N` to go to the previous result. Instead of `/`, you can use `?` to search backwards!

This is very useful when dealing with large manual pages where you need to quickly find specific information about options or arguments.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Open the manual page for the challenge:

```bash
   man challenge
```

3. Search for keywords related to the flag. Press `/` and type a search term:

```
   /flag
```

(Press Enter after typing)

4. Navigate through search results:

   - Press `n` to go to the next occurrence
   - Press `N` to go to the previous occurrence
   - Keep searching until you find the option that gives the flag
5. Once you find the correct option, press `q` to quit the manual
6. Run the challenge with the discovered option:

```bash
   /challenge/challenge [DISCOVERED_OPTION]
```

(Replace `[DISCOVERED_OPTION]` with the actual option found)

### Result

By using the search functionality (`/`) within the man page and navigating through results with `n` and `N`, the flag-providing option is quickly located without having to scroll through the entire manual. Running the challenge with this option successfully retrieves the flag. This demonstrates efficient navigation and search techniques within manual pages.



## Helpful Programs

### Task

Use the `--help` argument to read a program's built-in documentation and discover how to get the flag.

### Concept

Some programs don't have a man page, but might tell you how to run them if invoked with a special argument. Usually, this argument is `--help`, but it can often be `-h` or, in rare cases, `-?`, `help`, or other esoteric values like `/?` (though that latter is more frequently encountered on Windows).

In this level, you will practice reading a program's documentation with `--help`. Try it out!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Run the challenge program with the `--help` argument:

```bash
   /challenge/challenge --help
```

3. The help output shows:

```
   usage: a challenge to make you ask for help [-h] [--fortune] [-v] [-g GIVE_THE_FLAG] [-p]
   
   optional arguments:
     -h, --help            show this help message and exit
     --fortune             read your fortune
     -v, --version         get the version number
     -g GIVE_THE_FLAG, --give-the-flag GIVE_THE_FLAG
                           get the flag, if given the correct value
     -p, --print-value     print the value that will cause the -g option to give you the flag
```

4. Use the `-p` option to print the required value:

```bash
   /challenge/challenge -p
```

5. Use the `-g` option with the value obtained from step 4:

```bash
   /challenge/challenge -g [VALUE]
```

(Replace `[VALUE]` with the actual value printed by the `-p` option)

### Result

By reading the `--help` output, we discovered that the `-p` option prints the value needed for the `-g` option. First running with `-p` to get the required value, then running with `-g` and that value successfully retrieves the flag. This demonstrates how programs provide self-documentation through `--help` and how to use that information to solve multi-step challenges.


## Searching For Manuals

### Task

Search the man page database to find a randomly-named manpage for the challenge, then use it to discover the flag option.

### Concept

This level is tricky: it hides the manpage for the challenge by randomizing its name. Luckily, all of the manpages are gathered in a searchable database, so you'll be able to search the man page database to find the hidden challenge man page! To figure out how to search for the right manpage, read the `man` page manpage by doing: `man man`!

**HINT 1**: `man man` teaches you advanced usage of the `man` command itself, and you must use this knowledge to figure out how to search for the hidden manpage that will tell you how to use `/challenge/challenge`.

**HINT 2**: Though the manpage is randomly named, you still actually use `/challenge/challenge` to get the flag!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Read the manual page for `man` to learn how to search:

```bash
   man man
```

Look for options related to searching (hint: `-k` option searches man page descriptions)

3. Search the man page database for entries related to "challenge":

```bash
   man -k challenge
```

This will list all manpages with "challenge" in their name or description

4. Identify the randomly-named manpage from the search results (it will have a random name but relate to the challenge)
5. Read the discovered manpage:

```bash
   man [random-manpage-name]
```

(Replace `[random-manpage-name]` with the actual name found)

6. Find the option that gives the flag in this manpage
7. Run the challenge with the discovered option:

```bash
   /challenge/challenge [DISCOVERED_OPTION]
```

### Result

By using `man -k` to search the man page database, the randomly-named challenge manpage is located. Reading this manpage reveals the correct option to use. Running `/challenge/challenge` with this option successfully retrieves the flag. This demonstrates how to search for and discover manpages when their names are not immediately known.



## Help for Builtins

### Task

Use the `help` builtin command to get documentation for the `challenge` builtin and discover how to get the flag.

### Concept

Some commands, rather than being programs with man pages and help options, are built into the shell itself. These are called *builtins*. Builtins are invoked just like commands, but the shell handles them internally instead of launching other programs. You can get a list of shell builtins by running the builtin `help`, as so:

```bash
hacker@dojo:~$ help
```

You can get help on a specific one by passing it to the `help` builtin. Let's look at a builtin that we've already used earlier, `cd`:

```bash
hacker@dojo:~$ help cd
cd: cd [-L|[-P [-e]]] [dir]
    Change the shell working directory.
  
    Change the current directory to DIR.  The default DIR is the value of the
    HOME shell variable.
...
```

Some good information! In this challenge, we'll practice using `help` to look up help for builtins. This challenge's `challenge` command is a shell builtin, rather than a program. Like before, you need to lookup its help to figure out the secret value to pass to it!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Use the `help` builtin to get documentation for the `challenge` builtin:

```bash
   help challenge
```

3. Read the help output carefully to find:

   - The syntax/usage of the `challenge` builtin
   - The secret argument or value needed to get the flag
4. Run the `challenge` builtin with the discovered argument:

```bash
   challenge [DISCOVERED_ARGUMENT]
```

(Replace `[DISCOVERED_ARGUMENT]` with the actual argument found in the help output)

### Result

By using the `help` builtin to look up documentation for the `challenge` builtin, the correct argument is discovered. Running the builtin with this argument successfully retrieves the flag. This demonstrates how to get help for shell builtins, which don't have traditional man pages or `--help` options since they're part of the shell itself.




## Matching with *

### Task

Use the `*` glob wildcard to change directory to `/challenge` with an argument of at most 4 characters, then run the challenge.

### Concept

The first glob we'll learn is `*`. When it encounters a `*` character in any argument, the shell will treat it as a "wildcard" and try to replace that argument with any files that match the pattern. It's easier to show you than explain:

```bash
hacker@dojo:~$ touch file_a
hacker@dojo:~$ touch file_b
hacker@dojo:~$ touch file_c
hacker@dojo:~$ ls
file_a  file_b  file_c
hacker@dojo:~$ echo Look: file_*
Look: file_a file_b file_c
```

Of course, though in this case, the glob resulted in multiple arguments, it can just as simply match only one. For example:

```bash
hacker@dojo:~$ touch file_a
hacker@dojo:~$ ls
file_a
hacker@dojo:~$ echo Look: file_*
Look: file_a
```

When zero files are matched, by default, the shell leaves the glob unchanged:

```bash
hacker@dojo:~$ touch file_a
hacker@dojo:~$ ls
file_a
hacker@dojo:~$ echo Look: nope_*
Look: nope_*
```

The `*` matches any part of the filename except for `/` or a leading `.` character. For example:

```bash
hacker@dojo:~$ echo ONE: /ho*/*ck*
ONE: /home/hacker
hacker@dojo:~$ echo TWO: /*/hacker
TWO: /home/hacker
hacker@dojo:~$ echo THREE: ../*
THREE: ../hacker
```

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Change directory to `/challenge` using a glob pattern with at most 4 characters:

```bash
   cd /ch*
```

This uses only 4 characters (`/ch*`) and the `*` wildcard expands to match `/challenge`

3. Run the challenge program:

```bash
   /challenge/run
```

### Result

By using the glob pattern `/ch*` to match `/challenge`, we successfully navigate to the directory with an argument of only 4 characters. Running `/challenge/run` then retrieves the flag. This demonstrates how the `*` wildcard can be used to shorten command arguments while still matching the intended paths.




## Matching with ?

### Task

Use the `?` glob wildcard to change directory to `/challenge`, replacing specific characters with `?`, then run the challenge.

### Concept

Next, let's learn about `?`. When it encounters a `?` character in any argument, the shell will treat it as a single-character wildcard. This works like `*`, but only matches *one* character. For example:

```bash
hacker@dojo:~$ touch file_a
hacker@dojo:~$ touch file_b
hacker@dojo:~$ touch file_cc
hacker@dojo:~$ ls
file_a  file_b  file_cc
hacker@dojo:~$ echo Look: file_?
Look: file_a file_b
hacker@dojo:~$ echo Look: file_??
Look: file_cc
```

Now, practice this yourself! Starting from your home directory, change your directory to `/challenge`, but use the `?` character instead of `c` and `l` in the argument to `cd`! Once you're there, run `/challenge/run` for the flag!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Change directory to `/challenge` using `?` wildcards to replace `c` and `l`:

```bash
   cd /?ha??enge
```

This replaces:

- The `c` in "challenge" with `?`
- The `l` in "challenge" with `?`

Pattern: `/challenge` → `/?ha??enge`

3. Run the challenge program:

```bash
   /challenge/run
```

### Result

By using the `?` wildcard to match single characters, we successfully navigate to `/challenge` by replacing the `c` and `l` characters with `?`. Running `/challenge/run` then retrieves the flag. This demonstrates how the `?` wildcard matches exactly one character, unlike `*` which matches any number of characters.



## Matching paths with []

### Task

Use bracket globbing with absolute paths to match specific files and pass them as arguments to the challenge program.

### Concept

Globbing happens on a *path* basis, so you can expand entire paths with your globbed arguments. For example:

```bash
hacker@dojo:~$ touch file_a
hacker@dojo:~$ touch file_b
hacker@dojo:~$ touch file_c
hacker@dojo:~$ ls
file_a  file_b  file_c
hacker@dojo:~$ echo Look: /home/hacker/file_[ab]
Look: /home/hacker/file_a /home/hacker/file_b
```

Now it's your turn. Once more, we've placed a bunch of files in `/challenge/files`. Starting from your home directory, run `/challenge/run` with a single argument that bracket-globs into the absolute paths to the `file_b`, `file_a`, `file_s`, and `file_h` files!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. From the home directory, run the challenge with an absolute path bracket glob:

```bash
   /challenge/run /challenge/files/file_[bash]
```

This glob pattern will expand to the absolute paths:

- `/challenge/files/file_b`
- `/challenge/files/file_a`
- `/challenge/files/file_s`
- `/challenge/files/file_h`

### Result

By using the bracket glob with an absolute path `/challenge/files/file_[bash]`, the shell expands this to the full absolute paths of all four required files. The program receives these as arguments, verifies them, and provides the flag. This demonstrates how globbing works with complete paths, not just filenames, allowing you to match files anywhere in the filesystem without changing directories.

## Matching with []

### Task

Use the `[]` bracket glob to match specific files and pass them as arguments to the challenge program.

### Concept

Next, we will cover `[]`. The square brackets are, essentially, a limited form of `?`, in that instead of matching any character, `[]` is a wildcard for some subset of potential characters, specified within the brackets. For example, `[pwn]` will match the character `p`, `w`, or `n`. For example:

```bash
hacker@dojo:~$ touch file_a
hacker@dojo:~$ touch file_b
hacker@dojo:~$ touch file_c
hacker@dojo:~$ ls
file_a  file_b  file_c
hacker@dojo:~$ echo Look: file_[ab]
Look: file_a file_b
```

Try it here! We've placed a bunch of files in `/challenge/files`. Change your working directory to `/challenge/files` and run `/challenge/run` with a single argument that bracket-globs into `file_b`, `file_a`, `file_s`, and `file_h`!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Change directory to `/challenge/files`:

```bash
   cd /challenge/files
```

3. Run the challenge with a bracket glob that matches files with `b`, `a`, `s`, and `h`:

```bash
   /challenge/run file_[bash]
```

This glob pattern `file_[bash]` will match:

- `file_b` (matches 'b')
- `file_a` (matches 'a')
- `file_s` (matches 's')
- `file_h` (matches 'h')

### Result

By using the bracket glob `file_[bash]`, the shell expands this to match all four required files (`file_b`, `file_a`, `file_s`, `file_h`) and passes them as arguments to `/challenge/run`. The program verifies the correct files were matched and provides the flag. This demonstrates how `[]` can match specific characters from a defined set, providing more precise control than `*` or `?` wildcards.



## Matching paths with []

### Task

Use bracket globbing with absolute paths to match specific files and pass them as arguments to the challenge program.

### Concept

Globbing happens on a *path* basis, so you can expand entire paths with your globbed arguments. For example:

```bash
hacker@dojo:~$ touch file_a
hacker@dojo:~$ touch file_b
hacker@dojo:~$ touch file_c
hacker@dojo:~$ ls
file_a  file_b  file_c
hacker@dojo:~$ echo Look: /home/hacker/file_[ab]
Look: /home/hacker/file_a /home/hacker/file_b
```

Now it's your turn. Once more, we've placed a bunch of files in `/challenge/files`. Starting from your home directory, run `/challenge/run` with a single argument that bracket-globs into the absolute paths to the `file_b`, `file_a`, `file_s`, and `file_h` files!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. From the home directory, run the challenge with an absolute path bracket glob:

```bash
   /challenge/run /challenge/files/file_[bash]
```

This glob pattern will expand to the absolute paths:

- `/challenge/files/file_b`
- `/challenge/files/file_a`
- `/challenge/files/file_s`
- `/challenge/files/file_h`

### Result

By using the bracket glob with an absolute path `/challenge/files/file_[bash]`, the shell expands this to the full absolute paths of all four required files. The program receives these as arguments, verifies them, and provides the flag. This demonstrates how globbing works with complete paths, not just filenames, allowing you to match files anywhere in the filesystem without changing directories.



## Multiple Globs

### Task

Use multiple glob wildcards in a single argument to match files containing a specific letter.

### Concept

So far, you've specified one glob at a time, but you can do more! Bash supports the expansion of multiple globs in a single word. For example:

```bash
hacker@dojo:~$ cat /*fl*
pwn.college{YEAH}
hacker@dojo:~$
```

What happens above is that the shell looks for all files in `/` that start with *anything* (including nothing), then have an `f` and an `l`, and end in *anything* (including `ag`, which makes `flag`).

Now you try it. We put a few happy, but diversely-named files in `/challenge/files`. Go `cd` there and run `/challenge/run`, providing a single argument: a short (3 characters or less) globbed word with two `*` globs in it that covers every word that contains the letter `p`.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Change directory to `/challenge/files`:

```bash
   cd /challenge/files
```

3. Run the challenge with a glob pattern using two `*` wildcards that matches all files containing `p`:

```bash
   /challenge/run *p*
```

This pattern breakdown:

- First `*`: matches any characters before `p`
- `p`: matches the literal letter `p`
- Second `*`: matches any characters after `p`
- Total: 3 characters (`*p*`)

### Result

By using the glob pattern `*p*`, the shell expands this to match all files in the directory that contain the letter `p` anywhere in their name. The program receives all matching files as arguments and provides the flag. This demonstrates how multiple glob wildcards can be combined in a single pattern to create flexible matching rules.



## Mixing Globs

### Task

Combine different globbing techniques to create a short pattern that matches only specific files.

### Concept

Now, let's put the previous levels together! We put a few happy, but diversely-named files in `/challenge/files`. Go `cd` there and, using the globbing you've learned, write a single, short (6 characters or less) glob that (when passed as an argument to `/challenge/run`) will match only the files "challenging", "educational", and "pwning"!

**HINT**: Make sure to look at the names of the files in `/challenge/files`. Do you see any patterns that could help you make your glob?

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Change directory to `/challenge/files`:

```bash
   cd /challenge/files
```

3. List the files to observe patterns:

```bash
   ls
```

4. Analyze the target files:

   - "challenging" - starts with 'c'
   - "educational" - starts with 'e'
   - "pwning" - starts with 'p'
5. Create a glob pattern using bracket notation and wildcard:

```bash
   /challenge/run [cep]*
```

This pattern breakdown:

- `[cep]`: matches any file starting with 'c', 'e', or 'p'
- `*`: matches the rest of the filename
- Total: 6 characters (`[cep]*`)

### Result

By combining bracket notation `[]` with the wildcard `*`, the pattern `[cep]*` successfully matches only the three required files: "challenging", "educational", and "pwning". The program verifies the correct files were matched and provides the flag. This demonstrates how different globbing techniques can be mixed to create precise, efficient patterns for matching specific sets of files.



## Exclusionary Globbing

### Task

Use the `!` character inside brackets to create an exclusionary glob that matches files NOT starting with specific letters.

### Concept

Sometimes, you want to filter out files in a glob! Luckily, `[]` helps you do just this. If the first character in the brackets is a `!` or (in newer versions of bash) a `^`, the glob inverts, and that bracket instance matches characters that *aren't* listed. For example:

```bash
hacker@dojo:~$ touch file_a
hacker@dojo:~$ touch file_b
hacker@dojo:~$ touch file_c
hacker@dojo:~$ ls
file_a  file_b  file_c
hacker@dojo:~$ echo Look: file_[!ab]
Look: file_c
hacker@dojo:~$ echo Look: file_[^ab]
Look: file_c
hacker@dojo:~$ echo Look: file_[ab]
Look: file_a file_b
```

Armed with this knowledge, go forth to `/challenge/files` and run `/challenge/run` with all files that don't start with `p`, `w`, or `n`!

**NOTE**: The `!` character has a different special meaning in bash when it's not the first character of a `[]` glob, so keep that in mind if things stop making sense! `^` does not have this problem, but is also not compatible with older shells.

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Change directory to `/challenge/files`:

```bash
   cd /challenge/files
```

3. Run the challenge with an exclusionary glob pattern:

```bash
   /challenge/run [!pwn]*
```

This pattern breakdown:

- `[!pwn]`: matches any file that does NOT start with 'p', 'w', or 'n'
- `*`: matches the rest of the filename

Alternative (for newer bash versions):

```bash
   /challenge/run [^pwn]*
```

### Result

By using the exclusionary bracket notation `[!pwn]*`, the shell expands this to match all files in the directory that don't start with the letters 'p', 'w', or 'n'. The program receives these filtered files as arguments and provides the flag. This demonstrates how the `!` (or `^`) character inside brackets inverts the matching logic, allowing you to exclude specific characters from your glob patterns.



## Tab Completion

### Task

Use tab completion to auto-complete a tricky filename and read the flag file.

### Concept

As tempting as it might be, using `*` to shorten what must be typed on the commandline can lead to problems. Your glob might expand to unintended files, and you might not spot it until the `rm` command is already running! No one is safe from this style of error.

A safer alternative when you are trying to specify a specific target is **tab completion**. If you hit tab in the shell, it'll try to figure out what you're going to type and automatically complete it. Auto-completion is super useful, and this challenge will explore its use in specifying files.

This challenge has copied the flag into `/challenge/pwncollege`, and you can freely `cat` that file. But you can't type the filename: we used some serious trickery to make sure that you *must* tab-complete it. Try it out!

Example of tab completion in action:

```bash
hacker@dojo:~$ ls /challenge
DESCRIPTION.md  pwncollege
hacker@dojo:~$ cat /challenge/pwncollege
cat: /challenge/pwncollege: No such file or directory
hacker@dojo:~$ cat /challenge/pwn<TAB>
pwn.college{HECK YEAH}
hacker@dojo:~$
```

When you hit that tab key, the name will expand and you'll be able to read the file. Good luck!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Start typing the command to read the flag file:

```bash
   cat /challenge/pwn
```

3. Press the **Tab** key (without pressing Enter yet):

   - The shell will auto-complete the filename
   - The full filename with its tricky characters will appear
4. Press **Enter** to execute the command and read the flag

### Result

By using tab completion, the shell automatically fills in the complete filename including any special characters or variations that would be difficult to type manually. This reads the flag file successfully. This demonstrates how tab completion not only saves typing but also helps handle files with complex or hard-to-type names, while being safer than using wildcards that might match unintended files.



## Tab Completion on Commands

### Task

Use tab completion to auto-complete a command name and execute it to get the flag.

### Concept

Tab completion is for more than files! You can also tab-complete commands. This level has a command that starts with `pwncollege`, and it'll give you the flag. Type `pwncollege` and hit the tab key to auto-complete it!

**NOTE**: You can auto-complete any command, but be careful: callous auto-completes without double-checking the result can wreak havoc in your shell if you accidentally run the wrong commands!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Start typing the command name:

```bash
   pwncollege
```

3. Press the **Tab** key:

   - The shell will auto-complete to the full command name
   - The complete command will appear in your terminal
4. Press **Enter** to execute the command and retrieve the flag

### Result

By using tab completion on the command name, the shell automatically completes `pwncollege` to the full command name. Executing this command provides the flag. This demonstrates that tab completion works not only for files and directories but also for command names, making it easier to execute programs without typing their complete names.

**Important**: Always verify auto-completed commands before pressing Enter to avoid accidentally running unintended commands.


## Multiple Options for Tab Completion

### Task

Use tab completion when multiple matching files exist to navigate and find the flag file.

### Concept

Consider the following situation:

```bash
hacker@dojo:~$ ls
flag  flamingo  flowers
hacker@dojo:~$ cat f<TAB>
```

There are multiple options! What happens?

What happens varies based on the specific shell and its options. By default, `bash` will auto-expand until the first point when there are multiple options (in this case, `fl`). When you hit tab a *second* time, it'll print out those options. Other shells and configurations, instead, will cycle through the options.

This challenge has a `/challenge/files` directory with a bunch of files starting with `pwncollege`. Tab-complete from `/challenge/files/p` or so, and make your way to the flag!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Navigate to the challenge files directory:

```bash
   cd /challenge/files/
```

3. Start typing to read a file starting with "p":

```bash
   cat p
```

4. Press **Tab** once:

   - The shell will auto-complete to the common prefix of all matching files
   - Likely expands to `pwncollege`
5. Press **Tab** again:

   - The shell will display all files that start with the current prefix
   - You'll see a list of files starting with `pwncollege`
6. Continue typing additional characters to narrow down the options, then press **Tab** again to complete
7. Repeat this process until you find and complete the filename that contains the flag
8. Press **Enter** to read the file

### Result

By using tab completion multiple times and progressively narrowing down the options, the correct flag file is identified and read successfully. This demonstrates how tab completion handles multiple matching files by showing options and allowing progressive refinement of the target filename.




## Redirecting Output

### Task

Use output redirection to write the word "PWN" to a file named "COLLEGE".

### Concept

First, let's look at redirecting stdout to files. You can accomplish this with the `>` character, as so:

```bash
hacker@dojo:~$ echo hi > asdf
```

This will redirect the output of `echo hi` (which will be `hi`) to the file `asdf`. You can then use a program such as `cat` to output this file:

```bash
hacker@dojo:~$ cat asdf
hi
```

In this challenge, you must use this output redirection to write the word `PWN` (all uppercase) to the filename `COLLEGE` (all uppercase).

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Use `echo` with output redirection to write "PWN" to the file "COLLEGE":

```bash
   echo PWN > COLLEGE
```

3. The challenge will verify the file contents and provide the flag

### Result

By using the `>` redirection operator, the output of `echo PWN` is redirected to a file named `COLLEGE` instead of being printed to the terminal. The challenge verifies that the file contains the correct content and provides the flag. This demonstrates the basic concept of output redirection, which allows programs to write their output to files rather than displaying it on the screen.



## Redirecting More Output

### Task

Redirect the output of `/challenge/run` to a file named `myflag` to capture the flag.

### Concept

Aside from redirecting the output of `echo`, you can, of course, redirect the output of any command. In this level, `/challenge/run` will once more give you a flag, but *only* if you redirect its output to the file `myflag`. Your flag will, of course, end up in the `myflag` file!

You'll notice that `/challenge/run` will still happily print to your terminal, despite you redirecting stdout. That's because it communicates its instructions and feedback over standard error, and only prints the flag over standard out!

### Solution Steps

1. Connect to the dojo environment:

```bash
   ssh hacker@dojo.pwn.college
```

2. Run the challenge and redirect its output to `myflag`:

```bash
   /challenge/run > myflag
```

Note: You'll still see instructions printed to the terminal because they're sent to stderr, not stdout

3. Read the flag from the file:

```bash
   cat myflag
```

### Result

By redirecting the output of `/challenge/run` to the file `myflag`, the flag (which is sent to stdout) is captured in the file while instructions and feedback (sent to stderr) still appear on the terminal. Reading the `myflag` file with `cat` reveals the flag. This demonstrates how output redirection works with any command, and how programs can use different output streams (stdout vs stderr) for different types of information.
