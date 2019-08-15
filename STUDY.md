# STUDY


Let's do a dive into the world of Machine Learning Pipelines from the
perspective of _ML Infrastructure Architect_ or _Data Platform
Engineer_.


We will do that by going through a set of self-study materials and
using ML-Pipelines project as a reference.




## Chapter 1. Intro into the Domain

**Purpose of this intro**:

1. Start building up an intuition about the scope of the problem.
2. Identify for yourself if this field is something you are interested
   in.
3. Set up expectations about the complexity. All the other lessons
   would be easier to digest.

**What is Machine Learning?** Wikipedia defines it as:

> Machine learning (ML) is the scientific study of algorithms and
> statistical models that computer systems use to perform a specific
> task without using explicit instructions, relying on patterns and
> inference instead. It is seen as a subset of artificial intelligence.

All this sounds nice, but how does it apply to reality? Let's consider
an example domain of a company the specializes in producing and
selling robotic hands to supply warehouses and factories on a large
scale (we are obviously taking [Haddington
Dynamics](http://hdrobotic.com) as an inspiration).

Such a company would need to:

- manage research and development of robotic arm variants, end
  effectors (tools that are attached) and transmissions;
- manage supply chains to get hold of the needed parts in sufficient
  quality;
- maintain warehouse stock of the parts and spare parts;
- consider lead time (time it takes to get new part from
  the supplier);
- estimate potential demand for the goods sold based on the historical
  data;
- keep enough finished goods in stock to fulfill the demand;
- don't store too many goods in inventory, since things that sit idly
  in the warehouse are "frozen money" (that could quickly turn into
  the lost money).
  
  
Computers could help to manage all these business activities (mostly
when it comes to number crunching and estimating things).
  
In addition to managing the business side of the things, the company
would need to develop and maintain versions of the software that goes
into the robots. Some of that software would include ML models that
help robot to position itself properly (not a trivial task), react to
incoming forces or manage hardware degradation (even more tricky
tasks).

However, the domain complexity doesn't end up here. A growing company
would inherently have a number of departments, organizational
politics, a continuous stream of people that get hired, promoted or
leave the company. This creates need to manage data (not everybody
will have access to personal information of the customers) and
relations within the company (not everybody knows what everybody else
is doing). All that complicates operational landscape and makes it
more tricky to deploy ML solutions that affect the entire company.

_Zhamak Dehghani_ published an article on [How to Move Beyond a
Monolithic Data Lake to a Distributed Data
Mesh](https://martinfowler.com/articles/data-monolith-to-mesh.html). Glance
through this article, in order to gain a better understanding of the
scope of the problems within a data-driven company. You don't need to
understand everything at this point, just scan through the text.



A _ML Infrastructure Architect_ would be able to come to such a
company and lay out a realistic plan how to introduce Machine Learning
processes to this company and evolve them.

_Stephen Pimentel_ gave a 5 minute talk on [machine learning pipelines
at Apple](https://www.youtube.com/watch?v=16uU_Aaxp9Y). Watch it get
an idea of what steps does the ML Pipelines involve.



### Extra Reading

To gain more contextual understanding about building robotic hands
watch some [Haddington Dynamics Videos](http://hdrobotic.com/videos)
and read through these blog posts about replicating their technology:

1. [Child-friendly hobby](https://abdullin.com/child-friendly-hobby/)
2. [Dive into FPGA](https://abdullin.com/dive-into-fpga/)
3. [Running on a real FPGA](https://abdullin.com/running-on-a-real-fpga/)

  
## Chapter 2. Data Science is done in Python

_to be continued_
