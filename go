name: System tests
on: [push, pull_request]
jobs:
  tests:
    name: PHP ${{ matrix.php }}
    strategy:
      matrix:
        php: [8.2, 8.1, 8.0, 7.4, 7.3, 7.2, 7.1, 7.0]
    runs-on: ubuntu-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v3
      - name: Set up PHP
        uses: shivammathur/setup-php@v2
        with:
          php-version: ${{ matrix.php }}
          extensions: curl, gd, mbstring, zip
          ini-file: development
          coverage: none
          tools: none
      - name: Set up problem matcher
        run: echo "::add-matcher::${{ runner.tool_cache }}/php.json"
      - name: Set up test environment
        run: php yellow.php skip installation
      - name: Run tests
        run: php yellow.php build tests
