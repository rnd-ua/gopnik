include(ExternalProject)

#gtest
if(TESTS)
  set(GTEST_URL "http://googletest.googlecode.com/files/gtest-1.7.0.zip")
  if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/gtest-1.7.0.zip")
    set(GTEST_URL "${CMAKE_CURRENT_SOURCE_DIR}/gtest-1.7.0.zip")
  endif()

  ExternalProject_Add(gtest
    URL ${GTEST_URL}
    URL_MD5 "2d6ec8ccdf5c46b05ba54a9fd1d130d7"
    DOWNLOAD_DIR ${CMAKE_CURRENT_BINARY_DIR}
    SOURCE_DIR ${CMAKE_CURRENT_BINARY_DIR}/gtest-1.7.0
    BUILD_IN_SOURCE 1
    CONFIGURE_COMMAND "./configure"
    BUILD_COMMAND "make"
    INSTALL_COMMAND ""
  )

  set(GTEST_INCLUDE_DIR "${CMAKE_CURRENT_BINARY_DIR}/gtest-1.7.0/include" PARENT_SCOPE)
  set(GTEST_BOTH_LIBRARIES "${CMAKE_CURRENT_BINARY_DIR}/gtest-1.7.0/lib/.libs/libgtest.a" "${CMAKE_CURRENT_BINARY_DIR}/gtest-1.7.0/lib/.libs/libgtest.a" PARENT_SCOPE)
endif(TESTS)

if(BENCH)
  set(CELERO_URL "https://github.com/DigitalInBlue/Celero/archive/v1.0.5.tar.gz")
  if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/v1.0.5.tar.gz")
    set(CELERO_URL "${CMAKE_CURRENT_SOURCE_DIR}/v1.0.5.tar.gz")
  endif()

  ExternalProject_Add(celero
    URL ${CELERO_URL}
    URL_MD5 "e2e19b8ebe89591e3e75efa1c16ed7bd"
    DOWNLOAD_DIR ${CMAKE_CURRENT_BINARY_DIR}
    SOURCE_DIR ${CMAKE_CURRENT_BINARY_DIR}/Celero
    CMAKE_ARGS "-DCELERO_RUN_EXAMPLE_ON_BUILD=NO"
    BUILD_IN_SOURCE 1
    INSTALL_COMMAND ""
  )

  set(CELERO_INCLUDE_DIR "${CMAKE_CURRENT_BINARY_DIR}/Celero/include" PARENT_SCOPE)
  set(CELERO_LIBRARY "${CMAKE_CURRENT_BINARY_DIR}/Celero/libcelero.so" PARENT_SCOPE)
endif(BENCH)
