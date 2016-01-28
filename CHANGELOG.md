## 0.4.2
  * Returned striped hash implementation. It performs a bit worse on benchmark,
    but easy to undestand and prove correctness.

## 0.4.1
  * Implemented buffer pool. Server can accept messages of any size now.
  * Added unix socket support.
  * Added file logger.
  * Made some refactoring.
  * Improved benchmark.

## 0.4.0
  * Implemented lock-free hashmap. Huge performances improvements.

## 0.3.1
  * Small optimization of hasmap.getIndex method.

## 0.3.0
  * Implement basic server
  * Implement client
  * Sources structure refactoring

## 0.2.0
  * Implement Striped Hash algorithm
  * Implement FNV-1a hash function
  * Add concurrent benchmarks

## 0.1.1
  * Improve hashmap algorithm by substitution list by vector for collision resolution data structure.
  * Implement vector data structure.
  * Improve hashmap.getIndex method.

## 0.1.0
Initial version

  * Implement basic hashmap algorithm
  * Implement collections
  * Implement API
