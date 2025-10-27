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
