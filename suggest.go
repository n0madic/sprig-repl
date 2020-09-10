package main

import prompt "github.com/c-bata/go-prompt"

var suggest = []prompt.Suggest{
	{Text: "/help", Description: "Help for Sprig functions"},
	{Text: "/quit", Description: "Quit from REPL"},
	{Text: "abbrevboth", Description: "Abbreviate both sides"},
	{Text: "abbrev", Description: "Truncate a string with ellipses (...)"},
	{Text: "add1", Description: "To increment by 1"},
	{Text: "add", Description: "Sum numbers"},
	{Text: "adler32sum", Description: "Receives a string, and computes its Adler-32 checksum"},
	{Text: "ago", Description: "Returns duration from time.Now in seconds resolution"},
	{Text: "append", Description: "Append a new item to an existing list, creating a new list"},
	{Text: "atoi", Description: "Convert a string to an integer"},
	{Text: "b32dec", Description: "Decode with Base32"},
	{Text: "b32enc", Description: "Encode with Base32"},
	{Text: "b64dec", Description: "Decode with Base64"},
	{Text: "b64enc", Description: "Encode with Base64"},
	{Text: "base", Description: "Return the last element of a path"},
	{Text: "biggest", Description: "Return the largest of a series of integers"},
	{Text: "buildCustomCert", Description: "Allows customizing the certificate"},
	{Text: "camelcase", Description: "Convert string from snake_case to CamelCase"},
	{Text: "cat", Description: "Concatenates multiple strings together into one"},
	{Text: "ceil", Description: "Returns the greatest float value greater than or equal to input value"},
	{Text: "clean", Description: "Clean up a path"},
	{Text: "coalesce", Description: "Takes a list of values and returns the first non-empty one"},
	{Text: "compact", Description: "Accepts a list and removes entries with empty values"},
	{Text: "concat", Description: "Concatenate arbitrary number of lists into one"},
	{Text: "contains", Description: "Test to see if one string is contained inside of another"},
	{Text: "date", Description: "Formats a date"},
	{Text: "date_in_zone", Description: "Same as `date`, but with a timezone"},
	{Text: "dateInZone", Description: "Same as `date`, but with a timezone"},
	{Text: "date_modify", Description: "Takes a modification and a date and returns the timestamp"},
	{Text: "dateModify", Description: "Takes a modification and a date and returns the timestamp"},
	{Text: "decryptAES", Description: "Decode a base64 string encoded by the AES-256 CBC algorithm"},
	{Text: "deepCopy", Description: "Makes a deep copy of the value"},
	{Text: "deepEqual", Description: "Returns true if two values are deeply equal"},
	{Text: "default", Description: "Set a simple default value"},
	{Text: "derivePassword", Description: "Derive a specific password based on some shared “master password” constraints"},
	{Text: "dict", Description: "Creating dictionaries"},
	{Text: "dir", Description: "Return the directory, stripping the last part of the path"},
	{Text: "div", Description: "Perform integer division"},
	{Text: "duration", Description: "Formats a given amount of seconds as a time.Duration"},
	{Text: "durationRound", Description: "Rounds a given duration to the most significant unit"},
	{Text: "empty", Description: "Returns true if the given value is considered empty"},
	{Text: "encryptAES", Description: "Encrypts text with AES-256 CBC"},
	{Text: "env", Description: "Reads an environment variable"},
	{Text: "expandenv", Description: "Substitute environment variables in a string"},
	{Text: "ext", Description: "Return the file extension"},
	{Text: "fail", Description: "Unconditionally returns an empty string and an error with the specified text"},
	{Text: "first", Description: "Get the head item on a list"},
	{Text: "float64", Description: "Convert to a float64"},
	{Text: "floor", Description: "Returns the greatest float value less than or equal to input value"},
	{Text: "genCA", Description: "Generates a new, self-signed x509 certificate authority"},
	{Text: "genPrivateKey", Description: "Generates a new private key encoded into a PEM block"},
	{Text: "genSelfSignedCert", Description: "Generates a new, self-signed x509 certificate"},
	{Text: "genSignedCert", Description: "Generates a new, x509 certificate signed by the specified CA"},
	{Text: "get", Description: "Given a map and a key, get the value from the map"},
	{Text: "getHostByName", Description: "Receives a domain name and returns the ip address"},
	{Text: "has", Description: "Test to see if a list has a particular element"},
	{Text: "hasKey", Description: "Returns true if the given dict contains the given key"},
	{Text: "hasPrefix", Description: "Test whether a string has a given prefix"},
	{Text: "hasSuffix", Description: "Test whether a string has a given suffix"},
	{Text: "hello", Description: "Hello"},
	{Text: "htmlDate", Description: "Formats a date for inserting into an HTML date picker input field"},
	{Text: "htmlDateInZone", Description: "Same as `htmlDate`, but with a timezone"},
	{Text: "htpasswd", Description: "takes a `username` and `password` and generates a `bcrypt` hash of the password"},
	{Text: "indent", Description: "Indents every line in a given string to the specified indent width"},
	{Text: "initial", Description: "Return all but the last element"},
	{Text: "initials", Description: "Given multiple words, take the first letter of each word and combine"},
	{Text: "int64", Description: "Convert to an int64"},
	{Text: "int", Description: "Convert to an int at the system's width"},
	{Text: "isAbs", Description: "Check whether a file path is absolute"},
	{Text: "join", Description: "Join a list of strings into a single string, with the given separator"},
	{Text: "kebabcase", Description: "Convert string from camelCase to kebab-case"},
	{Text: "keys", Description: "Return a list of all of the keys in one or more dict types"},
	{Text: "kindIs", Description: "Verify that a value is a particular kind"},
	{Text: "kindOf", Description: "Returns the kind of an object"},
	{Text: "last", Description: "Get the last item on a list"},
	{Text: "list", Description: "Create a list"},
	{Text: "lower", Description: "Convert the entire string to lowercase"},
	{Text: "max", Description: "Return the largest of a series of integers"},
	{Text: "merge", Description: "Merge two or more dictionaries into one, giving precedence to the dest dictionary"},
	{Text: "mergeOverwrite", Description: "Merge two or more dictionaries into one, giving precedence from right to left"},
	{Text: "min", Description: "Return the smallest of a series of integers"},
	{Text: "mod", Description: "Modulo"},
	{Text: "mul", Description: "Multiply"},
	{Text: "mustAppend", Description: "Append a new item to an existing list, creating a new list"},
	{Text: "mustCompact", Description: "Accepts a list and removes entries with empty values"},
	{Text: "must_date_modify", Description: "Takes a modification and a date and returns the timestamp"},
	{Text: "mustDateModify", Description: "Takes a modification and a date and returns the timestamp"},
	{Text: "mustDeepCopy", Description: "Makes a deep copy of the value, returns an error when there is an error"},
	{Text: "mustFirst", Description: "Get the head item on a list"},
	{Text: "mustHas", Description: "Test to see if a list has a particular element"},
	{Text: "mustInitial", Description: "Return all but the last element"},
	{Text: "mustLast", Description: "Get the last item on a list"},
	{Text: "mustMerge", Description: "Merge two or more dictionaries into one, giving precedence to the dest dictionary"},
	{Text: "mustMergeOverwrite", Description: "Merge two or more dictionaries into one, giving precedence from right to left"},
	{Text: "mustPrepend", Description: "Push an alement onto the front of a list, creating a new list"},
	{Text: "mustPush", Description: "Append a new item to an existing list, creating a new list"},
	{Text: "mustRegexFindAll", Description: "Returns a slice of all matches of the regular expression in the input string"},
	{Text: "mustRegexFind", Description: "Return the first (left most) match of the regular expression in the input string"},
	{Text: "mustRegexMatch", Description: "Returns true if the input string contains any match of the regular expression"},
	{Text: "mustRegexReplaceAll", Description: "Replace matches of the Regexp with the replacement string replacement"},
	{Text: "mustRegexReplaceAllLiteral", Description: "Replace matches of the Regexp with the replacement string replacement"},
	{Text: "mustRegexSplit", Description: "Slices the input string into substrings separated by the expression and returns a slice"},
	{Text: "mustRest", Description: "Get the tail of the list (everything but the first item)"},
	{Text: "mustReverse", Description: "Produce a new list with the reversed elements of the given list"},
	{Text: "mustSlice", Description: "Get partial elements of a list"},
	{Text: "mustToDate", Description: "Converts a string to a date"},
	{Text: "mustToJson", Description: "Encodes an item into a JSON string"},
	{Text: "mustToPrettyJson", Description: "Encodes an item into a pretty (indented) JSON string"},
	{Text: "mustToRawJson", Description: "Encodes an item into JSON string with HTML characters unescaped"},
	{Text: "mustUniq", Description: "Generate a list with all of the duplicates removed"},
	{Text: "mustWithout", Description: "Filters items out of a list"},
	{Text: "nindent", Description: "Prepends a new line to the beginning of the string"},
	{Text: "nospace", Description: "Remove all whitespace from a string"},
	{Text: "now", Description: "The current date/time"},
	{Text: "omit", Description: "Returns a new dict with all the keys that do not match the given keys"},
	{Text: "pick", Description: "Selects just the given keys out of a dictionary, creating a new dict"},
	{Text: "pluck", Description: "Makes it possible to give one key and multiple maps, and get a list of all of the matches"},
	{Text: "plural", Description: "Pluralize a string"},
	{Text: "prepend", Description: "Push an alement onto the front of a list, creating a new list"},
	{Text: "push", Description: "Append a new item to an existing list, creating a new list"},
	{Text: "quote", Description: "Wrap a string in double quotes"},
	{Text: "randAlpha", Description: "Generate random strings uses a-zA-Z"},
	{Text: "randAlphaNum", Description: "Generate random strings uses 0-9a-zA-Z"},
	{Text: "randAscii", Description: "Generate random strings uses all printable ASCII characters"},
	{Text: "randNumeric", Description: "Generate random strings uses 0-9"},
	{Text: "regexFindAll", Description: "Returns a slice of all matches of the regular expression in the input string"},
	{Text: "regexFind", Description: "Return the first (left most) match of the regular expression in the input string"},
	{Text: "regexMatch", Description: "Returns true if the input string contains any match of the regular expression"},
	{Text: "regexReplaceAll", Description: "Replace matches of the Regexp with the replacement string replacement"},
	{Text: "regexReplaceAllLiteral", Description: "Replace matches of the Regexp with the replacement string replacement"},
	{Text: "regexSplit", Description: "Slices the input string into substrings separated by the expression and returns a slice"},
	{Text: "repeat", Description: "Repeat a string multiple times"},
	{Text: "replace", Description: "Perform simple string replacement"},
	{Text: "rest", Description: "Get the tail of the list (everything but the first item)"},
	{Text: "reverse", Description: "Produce a new list with the reversed elements of the given list"},
	{Text: "round", Description: "Returns a float value with the remainder rounded to the given number to digits after the decimal point"},
	{Text: "semverCompare", Description: "A more robust semver comparison"},
	{Text: "semver", Description: "Parses a string into a Semantic Version"},
	{Text: "seq", Description: "Works like the bash `seq` command"},
	{Text: "set", Description: "Add a new key/value pair to a dictionary"},
	{Text: "sha1sum", Description: "Receives a string, and computes it’s SHA1 digest"},
	{Text: "sha256sum", Description: "Receives a string, and computes it’s SHA256 digest"},
	{Text: "shuffle", Description: "Shuffle a string"},
	{Text: "slice", Description: "Get partial elements of a list"},
	{Text: "snakecase", Description: "Convert string from camelCase to snake_case"},
	{Text: "sortAlpha", Description: "Sorts a list of strings into alphabetical order"},
	{Text: "split", Description: "Splits a string into a dict"},
	{Text: "splitList", Description: "Split a string into a list of strings"},
	{Text: "splitn", Description: "Splits the number of substrings into a dict"},
	{Text: "squote", Description: "Wrap a string in single quotes"},
	{Text: "sub", Description: "To subtract"},
	{Text: "substr", Description: "Get a substring from a string"},
	{Text: "swapcase", Description: "Swap the case of a string using a word based algorithm"},
	{Text: "ternary", Description: "Similar to the C ternary operator"},
	{Text: "title", Description: "Convert to title case"},
	{Text: "toDate", Description: "Converts a string to a date"},
	{Text: "toDecimal", Description: "Convert a unix octal to a int64"},
	{Text: "toJson", Description: "Encodes an item into a JSON string"},
	{Text: "toPrettyJson", Description: "Encodes an item into a pretty (indented) JSON string"},
	{Text: "toRawJson", Description: "Encodes an item into JSON string with HTML characters unescaped"},
	{Text: "toString", Description: "Convert to a string"},
	{Text: "toStrings", Description: "Convert a list, slice, or array to a list of strings"},
	{Text: "trimall", Description: "Remove given characters from the front or back of a string"},
	{Text: "trimAll", Description: "Remove given characters from the front or back of a string"},
	{Text: "trim", Description: "Removes space from either side of a string"},
	{Text: "trimPrefix", Description: "Trim just the prefix from a string"},
	{Text: "trimSuffix", Description: "Trim just the suffix from a string"},
	{Text: "trunc", Description: "Truncate a string (and add no suffix)"},
	{Text: "tuple", Description: "Creates a tuple"},
	{Text: "typeIs", Description: "Verify that a value is a particular type"},
	{Text: "typeIsLike", Description: "Works as `typeIs`, except that it also dereferences pointers"},
	{Text: "typeOf", Description: "Returns the underlying type of a value"},
	{Text: "uniq", Description: "Generate a list with all of the duplicates removed"},
	{Text: "unixEpoch", Description: "Returns the seconds since the unix epoch for a time.Time"},
	{Text: "unset", Description: "Given a map and a key, delete the key from the map"},
	{Text: "until", Description: "Builds a range of integers"},
	{Text: "untilStep", Description: "Generates a list of counting integers"},
	{Text: "untitle", Description: "Remove title casing"},
	{Text: "upper", Description: "Convert the entire string to uppercase"},
	{Text: "urlJoin", Description: "Joins map (produced by urlParse) to produce URL string"},
	{Text: "urlParse", Description: "Parses string for URL and produces dict with URL parts"},
	{Text: "uuidv4", Description: "Generate UUID v4 universally unique IDs"},
	{Text: "values", Description: "Returns a new list with all the values of the source dict"},
	{Text: "without", Description: "Filters items out of a list"},
	{Text: "wrap", Description: "Wrap text at a given column count"},
	{Text: "wrapWith", Description: "Works as `wrap`, but lets you specify the string to wrap with"},
}
