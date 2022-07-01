# Usecase

- Similar to map but one key can store multiple values:


### Example:

```
Input: Push ("key1", "1"), Push ("key1", "2"), Push ("key2", "3"), Get "key1" 
Output: ["1", "2"]
```

### How to implement?
- A key will map with an slice of value
- key: comparable
- value: slice of comparable