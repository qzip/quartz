// Generated from /home/ashish/old2tb/AB2022Dev/projects/ONDC/ondc-server/clarity/parser/Clarity.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ClarityLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, STRING=5, UTF8=6, WHITESPACE=7, NUMBER=8, 
		HEX=9, SYMBOL=10, BUILTIN=11, PRINCIPAL=12, LPAREN=13, RPAREN=14, DOT=15, 
		BlockComment=16, LineComment=17;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "STRING", "UTF8", "WHITESPACE", "NUMBER", 
			"HEX", "SYMBOL", "BUILTIN", "PRINCIPAL", "LPAREN", "RPAREN", "DOT", "SYMBOL_START", 
			"DIGIT", "BlockComment", "LineComment"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "'{'", "','", "'}'", null, null, null, null, null, null, 
			null, null, "'('", "')'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "STRING", "UTF8", "WHITESPACE", "NUMBER", 
			"HEX", "SYMBOL", "BUILTIN", "PRINCIPAL", "LPAREN", "RPAREN", "DOT", "BlockComment", 
			"LineComment"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ClarityLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Clarity.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\23\u0091\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\7\6\66\n\6\f\6\16\69\13\6\3\6\3\6\3\7\3\7\3\7\3\b\6\bA\n\b\r\b\16"+
		"\bB\3\b\3\b\3\t\5\tH\n\t\3\t\6\tK\n\t\r\t\16\tL\3\t\3\t\6\tQ\n\t\r\t\16"+
		"\tR\5\tU\n\t\3\t\5\tX\n\t\3\n\3\n\3\n\3\n\5\n^\n\n\3\13\3\13\3\13\7\13"+
		"c\n\13\f\13\16\13f\13\13\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\17\3\17\3"+
		"\20\3\20\3\21\5\21u\n\21\3\22\3\22\3\23\3\23\3\23\3\23\7\23}\n\23\f\23"+
		"\16\23\u0080\13\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\7\24\u008b"+
		"\n\24\f\24\16\24\u008e\13\24\3\24\3\24\3~\2\25\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\2#\2%\22\'\23\3"+
		"\2\t\4\2$$^^\5\2\13\f\17\17\"\"\4\2--//\4\2CHch\4\2##AA\7\2,-/\61C\\a"+
		"ac|\4\2\f\f\17\17\2\u009c\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
		"\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2"+
		"\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3"+
		"\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\3)\3\2\2\2\5+\3\2\2\2\7-\3\2\2\2\t/\3\2"+
		"\2\2\13\61\3\2\2\2\r<\3\2\2\2\17@\3\2\2\2\21W\3\2\2\2\23Y\3\2\2\2\25_"+
		"\3\2\2\2\27g\3\2\2\2\31j\3\2\2\2\33m\3\2\2\2\35o\3\2\2\2\37q\3\2\2\2!"+
		"t\3\2\2\2#v\3\2\2\2%x\3\2\2\2\'\u0086\3\2\2\2)*\7<\2\2*\4\3\2\2\2+,\7"+
		"}\2\2,\6\3\2\2\2-.\7.\2\2.\b\3\2\2\2/\60\7\177\2\2\60\n\3\2\2\2\61\67"+
		"\7$\2\2\62\63\7^\2\2\63\66\13\2\2\2\64\66\n\2\2\2\65\62\3\2\2\2\65\64"+
		"\3\2\2\2\669\3\2\2\2\67\65\3\2\2\2\678\3\2\2\28:\3\2\2\29\67\3\2\2\2:"+
		";\7$\2\2;\f\3\2\2\2<=\7w\2\2=>\5\13\6\2>\16\3\2\2\2?A\t\3\2\2@?\3\2\2"+
		"\2AB\3\2\2\2B@\3\2\2\2BC\3\2\2\2CD\3\2\2\2DE\b\b\2\2E\20\3\2\2\2FH\t\4"+
		"\2\2GF\3\2\2\2GH\3\2\2\2HJ\3\2\2\2IK\5#\22\2JI\3\2\2\2KL\3\2\2\2LJ\3\2"+
		"\2\2LM\3\2\2\2MT\3\2\2\2NP\7\60\2\2OQ\5#\22\2PO\3\2\2\2QR\3\2\2\2RP\3"+
		"\2\2\2RS\3\2\2\2SU\3\2\2\2TN\3\2\2\2TU\3\2\2\2UX\3\2\2\2VX\5\23\n\2WG"+
		"\3\2\2\2WV\3\2\2\2X\22\3\2\2\2YZ\7\62\2\2Z]\7z\2\2[^\5#\22\2\\^\t\5\2"+
		"\2][\3\2\2\2]\\\3\2\2\2]^\3\2\2\2^\24\3\2\2\2_d\5!\21\2`c\5!\21\2ac\5"+
		"#\22\2b`\3\2\2\2ba\3\2\2\2cf\3\2\2\2db\3\2\2\2de\3\2\2\2e\26\3\2\2\2f"+
		"d\3\2\2\2gh\5\25\13\2hi\t\6\2\2i\30\3\2\2\2jk\7)\2\2kl\5\25\13\2l\32\3"+
		"\2\2\2mn\7*\2\2n\34\3\2\2\2op\7+\2\2p\36\3\2\2\2qr\7\60\2\2r \3\2\2\2"+
		"su\t\7\2\2ts\3\2\2\2u\"\3\2\2\2vw\4\62;\2w$\3\2\2\2xy\7\61\2\2yz\7,\2"+
		"\2z~\3\2\2\2{}\13\2\2\2|{\3\2\2\2}\u0080\3\2\2\2~\177\3\2\2\2~|\3\2\2"+
		"\2\177\u0081\3\2\2\2\u0080~\3\2\2\2\u0081\u0082\7,\2\2\u0082\u0083\7\61"+
		"\2\2\u0083\u0084\3\2\2\2\u0084\u0085\b\23\2\2\u0085&\3\2\2\2\u0086\u0087"+
		"\7=\2\2\u0087\u0088\7=\2\2\u0088\u008c\3\2\2\2\u0089\u008b\n\b\2\2\u008a"+
		"\u0089\3\2\2\2\u008b\u008e\3\2\2\2\u008c\u008a\3\2\2\2\u008c\u008d\3\2"+
		"\2\2\u008d\u008f\3\2\2\2\u008e\u008c\3\2\2\2\u008f\u0090\b\24\2\2\u0090"+
		"(\3\2\2\2\21\2\65\67BGLRTW]bdt~\u008c\3\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}